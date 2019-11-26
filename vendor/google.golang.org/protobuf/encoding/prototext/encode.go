// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototext

import (
	"fmt"
	"sort"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/encoding/text"
	"google.golang.org/protobuf/internal/encoding/wire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/fieldnum"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/mapsort"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/proto"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Marshal writes the given proto.Message in textproto format using default
// options. Do not depend on the output being stable. It may change over time
// across different versions of the program.
func Marshal(m proto.Message) ([]byte, error) {
	return MarshalOptions{}.Marshal(m)
}

// MarshalOptions is a configurable text format marshaler.
type MarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// AllowPartial allows messages that have missing required fields to marshal
	// without returning an error. If AllowPartial is false (the default),
	// Marshal will return error if there are any missing required fields.
	AllowPartial bool

	// If Indent is a non-empty string, it causes entries for a Message to be
	// preceded by the indent and trailed by a newline. Indent can only be
	// composed of space or tab characters.
	Indent string

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Marshal writes the given proto.Message in textproto format using options in
// MarshalOptions object. Do not depend on the output being stable. It may
// change over time across different versions of the program.
func (o MarshalOptions) Marshal(m proto.Message) ([]byte, error) {
	if o.Resolver == nil {
		o.Resolver = protoregistry.GlobalTypes
	}

	v, err := o.marshalMessage(m.ProtoReflect())
	if err != nil {
		return nil, err
	}

	delims := [2]byte{'{', '}'}
	const outputASCII = false
	b, err := text.Marshal(v, o.Indent, delims, outputASCII)
	if err != nil {
		return nil, err
	}
	if o.AllowPartial {
		return b, nil
	}
	return b, proto.IsInitialized(m)
}

// marshalMessage converts a protoreflect.Message to a text.Value.
func (o MarshalOptions) marshalMessage(m pref.Message) (text.Value, error) {
	messageDesc := m.Descriptor()
	if !flags.ProtoLegacy && messageset.IsMessageSet(messageDesc) {
		return text.Value{}, errors.New("no support for proto1 MessageSets")
	}

	// Handle Any expansion.
	if messageDesc.FullName() == "google.protobuf.Any" {
		if msg, err := o.marshalAny(m); err == nil {
			// Return as is if no error.
			return msg, nil
		}
		// Otherwise continue on to marshal Any as a regular message.
	}

	// Handle known fields.
	var msgFields [][2]text.Value
	fieldDescs := messageDesc.Fields()
	size := fieldDescs.Len()
	for i := 0; i < size; i++ {
		fd := fieldDescs.Get(i)
		if !m.Has(fd) {
			continue
		}

		name := text.ValueOf(fd.Name())
		// Use type name for group field name.
		if fd.Kind() == pref.GroupKind {
			name = text.ValueOf(fd.Message().Name())
		}
		pval := m.Get(fd)
		var err error
		msgFields, err = o.appendField(msgFields, name, pval, fd)
		if err != nil {
			return text.Value{}, err
		}
	}

	// Handle extensions.
	var err error
	msgFields, err = o.appendExtensions(msgFields, m)
	if err != nil {
		return text.Value{}, err
	}

	// Handle unknown fields.
	// TODO: Provide option to exclude or include unknown fields.
	msgFields = appendUnknown(msgFields, m.GetUnknown())

	return text.ValueOf(msgFields), nil
}

// appendField marshals a protoreflect.Value and appends it to the given [][2]text.Value.
func (o MarshalOptions) appendField(msgFields [][2]text.Value, name text.Value, pval pref.Value, fd pref.FieldDescriptor) ([][2]text.Value, error) {
	switch {
	case fd.IsList():
		items, err := o.marshalList(pval.List(), fd)
		if err != nil {
			return msgFields, err
		}

		for _, item := range items {
			msgFields = append(msgFields, [2]text.Value{name, item})
		}
	case fd.IsMap():
		items, err := o.marshalMap(pval.Map(), fd)
		if err != nil {
			return msgFields, err
		}

		for _, item := range items {
			msgFields = append(msgFields, [2]text.Value{name, item})
		}
	default:
		tval, err := o.marshalSingular(pval, fd)
		if err != nil {
			return msgFields, err
		}
		msgFields = append(msgFields, [2]text.Value{name, tval})
	}

	return msgFields, nil
}

// marshalSingular converts a non-repeated field value to text.Value.
// This includes all scalar types, enums, messages, and groups.
func (o MarshalOptions) marshalSingular(val pref.Value, fd pref.FieldDescriptor) (text.Value, error) {
	kind := fd.Kind()
	switch kind {
	case pref.BoolKind,
		pref.Int32Kind, pref.Sint32Kind, pref.Uint32Kind,
		pref.Int64Kind, pref.Sint64Kind, pref.Uint64Kind,
		pref.Sfixed32Kind, pref.Fixed32Kind,
		pref.Sfixed64Kind, pref.Fixed64Kind,
		pref.FloatKind, pref.DoubleKind,
		pref.BytesKind:
		return text.ValueOf(val.Interface()), nil

	case pref.StringKind:
		s := val.String()
		if !utf8.ValidString(s) {
			return text.Value{}, errors.InvalidUTF8(string(fd.FullName()))
		}
		return text.ValueOf(s), nil

	case pref.EnumKind:
		num := val.Enum()
		if desc := fd.Enum().Values().ByNumber(num); desc != nil {
			return text.ValueOf(desc.Name()), nil
		}
		// Use numeric value if there is no enum description.
		return text.ValueOf(int32(num)), nil

	case pref.MessageKind, pref.GroupKind:
		return o.marshalMessage(val.Message())
	}

	panic(fmt.Sprintf("%v has unknown kind: %v", fd.FullName(), kind))
}

// marshalList converts a protoreflect.List to []text.Value.
func (o MarshalOptions) marshalList(list pref.List, fd pref.FieldDescriptor) ([]text.Value, error) {
	size := list.Len()
	values := make([]text.Value, 0, size)

	for i := 0; i < size; i++ {
		item := list.Get(i)
		val, err := o.marshalSingular(item, fd)
		if err != nil {
			// Return already marshaled values.
			return values, err
		}
		values = append(values, val)
	}

	return values, nil
}

var (
	mapKeyName   = text.ValueOf(pref.Name("key"))
	mapValueName = text.ValueOf(pref.Name("value"))
)

// marshalMap converts a protoreflect.Map to []text.Value.
func (o MarshalOptions) marshalMap(mmap pref.Map, fd pref.FieldDescriptor) ([]text.Value, error) {
	// values is a list of messages.
	values := make([]text.Value, 0, mmap.Len())

	var err error
	mapsort.Range(mmap, fd.MapKey().Kind(), func(key pref.MapKey, val pref.Value) bool {
		var keyTxtVal text.Value
		keyTxtVal, err = o.marshalSingular(key.Value(), fd.MapKey())
		if err != nil {
			return false
		}
		var valTxtVal text.Value
		valTxtVal, err = o.marshalSingular(val, fd.MapValue())
		if err != nil {
			return false
		}
		// Map entry (message) contains 2 fields, first field for key and second field for value.
		msg := text.ValueOf([][2]text.Value{
			{mapKeyName, keyTxtVal},
			{mapValueName, valTxtVal},
		})
		values = append(values, msg)
		err = nil
		return true
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

// appendExtensions marshals extension fields and appends them to the given [][2]text.Value.
func (o MarshalOptions) appendExtensions(msgFields [][2]text.Value, m pref.Message) ([][2]text.Value, error) {
	var err error
	var entries [][2]text.Value
	m.Range(func(fd pref.FieldDescriptor, v pref.Value) bool {
		if !fd.IsExtension() {
			return true
		}

		// For MessageSet extensions, the name used is the parent message.
		name := fd.FullName()
		if messageset.IsMessageSetExtension(fd) {
			name = name.Parent()
		}

		// Use string type to produce [name] format.
		tname := text.ValueOf(string(name))
		entries, err = o.appendField(entries, tname, v, fd)
		if err != nil {
			return false
		}
		err = nil
		return true
	})
	if err != nil {
		return msgFields, err
	}

	// Sort extensions lexicographically and append to output.
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i][0].String() < entries[j][0].String()
	})
	return append(msgFields, entries...), nil
}

// appendUnknown parses the given []byte and appends field(s) into the given fields slice.
// This function assumes proper encoding in the given []byte.
func appendUnknown(fields [][2]text.Value, b []byte) [][2]text.Value {
	for len(b) > 0 {
		var value interface{}
		num, wtype, n := wire.ConsumeTag(b)
		b = b[n:]

		switch wtype {
		case wire.VarintType:
			value, n = wire.ConsumeVarint(b)
		case wire.Fixed32Type:
			value, n = wire.ConsumeFixed32(b)
		case wire.Fixed64Type:
			value, n = wire.ConsumeFixed64(b)
		case wire.BytesType:
			value, n = wire.ConsumeBytes(b)
		case wire.StartGroupType:
			var v []byte
			v, n = wire.ConsumeGroup(num, b)
			var msg [][2]text.Value
			value = appendUnknown(msg, v)
		default:
			panic(fmt.Sprintf("error parsing unknown field wire type: %v", wtype))
		}

		fields = append(fields, [2]text.Value{text.ValueOf(uint32(num)), text.ValueOf(value)})
		b = b[n:]
	}
	return fields
}

// marshalAny converts a google.protobuf.Any protoreflect.Message to a text.Value.
func (o MarshalOptions) marshalAny(m pref.Message) (text.Value, error) {
	fds := m.Descriptor().Fields()
	fdType := fds.ByNumber(fieldnum.Any_TypeUrl)
	fdValue := fds.ByNumber(fieldnum.Any_Value)

	typeURL := m.Get(fdType).String()
	value := m.Get(fdValue)

	emt, err := o.Resolver.FindMessageByURL(typeURL)
	if err != nil {
		return text.Value{}, err
	}
	em := emt.New().Interface()
	err = proto.UnmarshalOptions{
		AllowPartial: true,
		Resolver:     o.Resolver,
	}.Unmarshal(value.Bytes(), em)
	if err != nil {
		return text.Value{}, err
	}

	msg, err := o.marshalMessage(em.ProtoReflect())
	if err != nil {
		return text.Value{}, err
	}
	// Expanded Any field value contains only a single field with the type_url field value as the
	// field name in [] and a text marshaled field value of the embedded message.
	msgFields := [][2]text.Value{
		{
			text.ValueOf(typeURL),
			msg,
		},
	}
	return text.ValueOf(msgFields), nil
}
