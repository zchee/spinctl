// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoregistry provides data structures to register and lookup
// protobuf descriptor types.
//
// The Files registry contains file descriptors and provides the ability
// to iterate over the files or lookup a specific descriptor within the files.
// Files only contains protobuf descriptors and has no understanding of Go
// type information that may be associated with each descriptor.
//
// The Types registry contains descriptor types for which there is a known
// Go type associated with that descriptor. It provides the ability to iterate
// over the registered types or lookup a type by name.
package protoregistry

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ignoreConflict reports whether to ignore a registration conflict
// given the descriptor being registered and the error.
// It is a variable so that the behavior is easily overridden in another file.
var ignoreConflict = func(d protoreflect.Descriptor, err error) bool {
	log.Printf(""+
		"WARNING: %v\n"+
		"A future release will panic on registration conflicts.\n"+
		// TODO: Add a URL pointing to documentation on how to resolve conflicts.
		"\n", err)
	return true
}

// GlobalFiles is a global registry of file descriptors.
var GlobalFiles *Files = new(Files)

// GlobalTypes is the registry used by default for type lookups
// unless a local registry is provided by the user.
var GlobalTypes *Types = new(Types)

// NotFound is a sentinel error value to indicate that the type was not found.
var NotFound = errors.New("not found")

// Files is a registry for looking up or iterating over files and the
// descriptors contained within them.
// The Find and Range methods are safe for concurrent use.
type Files struct {
	// The map of descsByName contains:
	//	EnumDescriptor
	//	EnumValueDescriptor
	//	MessageDescriptor
	//	ExtensionDescriptor
	//	ServiceDescriptor
	//	*packageDescriptor
	//
	// Note that files are stored as a slice, since a package may contain
	// multiple files. Only top-level declarations are registered.
	// Note that enum values are in the top-level since that are in the same
	// scope as the parent enum.
	descsByName map[protoreflect.FullName]interface{}
	filesByPath map[string]protoreflect.FileDescriptor
}

type packageDescriptor struct {
	files []protoreflect.FileDescriptor
}

// NewFiles returns a registry initialized with the provided set of files.
// Files with a namespace conflict with an pre-existing file are not registered.
func NewFiles(files ...protoreflect.FileDescriptor) *Files {
	r := new(Files)
	r.Register(files...) // ignore errors; first takes precedence
	return r
}

// Register registers the provided list of file descriptors.
//
// If any descriptor within a file conflicts with the descriptor of any
// previously registered file (e.g., two enums with the same full name),
// then that file is not registered and an error is returned.
//
// It is permitted for multiple files to have the same file path.
func (r *Files) Register(files ...protoreflect.FileDescriptor) error {
	if r.descsByName == nil {
		r.descsByName = map[protoreflect.FullName]interface{}{
			"": &packageDescriptor{},
		}
		r.filesByPath = make(map[string]protoreflect.FileDescriptor)
	}
	var firstErr error
	for _, file := range files {
		if err := r.registerFile(file); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
func (r *Files) registerFile(fd protoreflect.FileDescriptor) error {
	path := fd.Path()
	if prev := r.filesByPath[path]; prev != nil {
		err := errors.New("file %q is already registered", fd.Path())
		err = amendErrorWithCaller(err, prev, fd)
		if r == GlobalFiles && ignoreConflict(fd, err) {
			err = nil
		}
		return err
	}

	for name := fd.Package(); name != ""; name = name.Parent() {
		switch prev := r.descsByName[name]; prev.(type) {
		case nil, *packageDescriptor:
		default:
			err := errors.New("file %q has a package name conflict over %v", fd.Path(), name)
			err = amendErrorWithCaller(err, prev, fd)
			if r == GlobalFiles && ignoreConflict(fd, err) {
				err = nil
			}
			return err
		}
	}
	var err error
	var hasConflict bool
	rangeTopLevelDescriptors(fd, func(d protoreflect.Descriptor) {
		if prev := r.descsByName[d.FullName()]; prev != nil {
			hasConflict = true
			err = errors.New("file %q has a name conflict over %v", fd.Path(), d.FullName())
			err = amendErrorWithCaller(err, prev, fd)
			if r == GlobalFiles && ignoreConflict(d, err) {
				err = nil
			}
		}
	})
	if hasConflict {
		return err
	}

	for name := fd.Package(); name != ""; name = name.Parent() {
		if r.descsByName[name] == nil {
			r.descsByName[name] = &packageDescriptor{}
		}
	}
	p := r.descsByName[fd.Package()].(*packageDescriptor)
	p.files = append(p.files, fd)
	rangeTopLevelDescriptors(fd, func(d protoreflect.Descriptor) {
		r.descsByName[d.FullName()] = d
	})
	r.filesByPath[path] = fd
	return nil
}

// FindDescriptorByName looks up a descriptor by the full name.
//
// This returns (nil, NotFound) if not found.
func (r *Files) FindDescriptorByName(name protoreflect.FullName) (protoreflect.Descriptor, error) {
	if r == nil {
		return nil, NotFound
	}
	prefix := name
	suffix := nameSuffix("")
	for prefix != "" {
		if d, ok := r.descsByName[prefix]; ok {
			switch d := d.(type) {
			case protoreflect.EnumDescriptor:
				if d.FullName() == name {
					return d, nil
				}
			case protoreflect.EnumValueDescriptor:
				if d.FullName() == name {
					return d, nil
				}
			case protoreflect.MessageDescriptor:
				if d.FullName() == name {
					return d, nil
				}
				if d := findDescriptorInMessage(d, suffix); d != nil && d.FullName() == name {
					return d, nil
				}
			case protoreflect.ExtensionDescriptor:
				if d.FullName() == name {
					return d, nil
				}
			case protoreflect.ServiceDescriptor:
				if d.FullName() == name {
					return d, nil
				}
				if d := d.Methods().ByName(suffix.Pop()); d != nil && d.FullName() == name {
					return d, nil
				}
			}
			return nil, NotFound
		}
		prefix = prefix.Parent()
		suffix = nameSuffix(name[len(prefix)+len("."):])
	}
	return nil, NotFound
}

func findDescriptorInMessage(md protoreflect.MessageDescriptor, suffix nameSuffix) protoreflect.Descriptor {
	name := suffix.Pop()
	if suffix == "" {
		if ed := md.Enums().ByName(name); ed != nil {
			return ed
		}
		for i := md.Enums().Len() - 1; i >= 0; i-- {
			if vd := md.Enums().Get(i).Values().ByName(name); vd != nil {
				return vd
			}
		}
		if xd := md.Extensions().ByName(name); xd != nil {
			return xd
		}
		if fd := md.Fields().ByName(name); fd != nil {
			return fd
		}
		if od := md.Oneofs().ByName(name); od != nil {
			return od
		}
	}
	if md := md.Messages().ByName(name); md != nil {
		if suffix == "" {
			return md
		}
		return findDescriptorInMessage(md, suffix)
	}
	return nil
}

type nameSuffix string

func (s *nameSuffix) Pop() (name protoreflect.Name) {
	if i := strings.IndexByte(string(*s), '.'); i >= 0 {
		name, *s = protoreflect.Name((*s)[:i]), (*s)[i+1:]
	} else {
		name, *s = protoreflect.Name((*s)), ""
	}
	return name
}

// FindFileByPath looks up a file by the path.
//
// This returns (nil, NotFound) if not found.
func (r *Files) FindFileByPath(path string) (protoreflect.FileDescriptor, error) {
	if r == nil {
		return nil, NotFound
	}
	if fd, ok := r.filesByPath[path]; ok {
		return fd, nil
	}
	return nil, NotFound
}

// NumFiles reports the number of registered files.
func (r *Files) NumFiles() int {
	if r == nil {
		return 0
	}
	return len(r.filesByPath)
}

// RangeFiles iterates over all registered files.
// The iteration order is undefined.
func (r *Files) RangeFiles(f func(protoreflect.FileDescriptor) bool) {
	if r == nil {
		return
	}
	for _, file := range r.filesByPath {
		if !f(file) {
			return
		}
	}
}

// NumFilesByPackage reports the number of registered files in a proto package.
func (r *Files) NumFilesByPackage(name protoreflect.FullName) int {
	if r == nil {
		return 0
	}
	p, ok := r.descsByName[name].(*packageDescriptor)
	if !ok {
		return 0
	}
	return len(p.files)
}

// RangeFilesByPackage iterates over all registered files in a give proto package.
// The iteration order is undefined.
func (r *Files) RangeFilesByPackage(name protoreflect.FullName, f func(protoreflect.FileDescriptor) bool) {
	if r == nil {
		return
	}
	p, ok := r.descsByName[name].(*packageDescriptor)
	if !ok {
		return
	}
	for _, file := range p.files {
		if !f(file) {
			return
		}
	}
}

// rangeTopLevelDescriptors iterates over all top-level descriptors in a file
// which will be directly entered into the registry.
func rangeTopLevelDescriptors(fd protoreflect.FileDescriptor, f func(protoreflect.Descriptor)) {
	eds := fd.Enums()
	for i := eds.Len() - 1; i >= 0; i-- {
		f(eds.Get(i))
		vds := eds.Get(i).Values()
		for i := vds.Len() - 1; i >= 0; i-- {
			f(vds.Get(i))
		}
	}
	mds := fd.Messages()
	for i := mds.Len() - 1; i >= 0; i-- {
		f(mds.Get(i))
	}
	xds := fd.Extensions()
	for i := xds.Len() - 1; i >= 0; i-- {
		f(xds.Get(i))
	}
	sds := fd.Services()
	for i := sds.Len() - 1; i >= 0; i-- {
		f(sds.Get(i))
	}
}

// Type is an interface satisfied by protoreflect.EnumType,
// protoreflect.MessageType, or protoreflect.ExtensionType.
type Type interface {
	GoType() reflect.Type
}

var (
	_ Type = protoreflect.EnumType(nil)
	_ Type = protoreflect.MessageType(nil)
	_ Type = protoreflect.ExtensionType(nil)
)

// MessageTypeResolver is an interface for looking up messages.
//
// A compliant implementation must deterministically return the same type
// if no error is encountered.
//
// The Types type implements this interface.
type MessageTypeResolver interface {
	// FindMessageByName looks up a message by its full name.
	// E.g., "google.protobuf.Any"
	//
	// This return (nil, NotFound) if not found.
	FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error)

	// FindMessageByURL looks up a message by a URL identifier.
	// See documentation on google.protobuf.Any.type_url for the URL format.
	//
	// This returns (nil, NotFound) if not found.
	FindMessageByURL(url string) (protoreflect.MessageType, error)
}

// ExtensionTypeResolver is an interface for looking up extensions.
//
// A compliant implementation must deterministically return the same type
// if no error is encountered.
//
// The Types type implements this interface.
type ExtensionTypeResolver interface {
	// FindExtensionByName looks up a extension field by the field's full name.
	// Note that this is the full name of the field as determined by
	// where the extension is declared and is unrelated to the full name of the
	// message being extended.
	//
	// This returns (nil, NotFound) if not found.
	FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)

	// FindExtensionByNumber looks up a extension field by the field number
	// within some parent message, identified by full name.
	//
	// This returns (nil, NotFound) if not found.
	FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
}

var (
	_ MessageTypeResolver   = (*Types)(nil)
	_ ExtensionTypeResolver = (*Types)(nil)
)

// Types is a registry for looking up or iterating over descriptor types.
// The Find and Range methods are safe for concurrent use.
type Types struct {
	// TODO: The syntax of the URL is ill-defined and the protobuf team recently
	// changed the documented semantics in a way that breaks prior usages.
	// I do not believe they can do this and need to sync up with the
	// protobuf team again to hash out what the proper syntax of the URL is.

	// TODO: Should we separate this out as a registry for each type?
	//
	// In Java, the extension and message registry are distinct classes.
	// Their extension registry has knowledge of distinct Java types,
	// while their message registry only contains descriptor information.
	//
	// In Go, we have always registered messages, enums, and extensions.
	// Messages and extensions are registered with Go information, while enums
	// are only registered with descriptor information. We cannot drop Go type
	// information for messages otherwise we would be unable to implement
	// portions of the v1 API such as ptypes.DynamicAny.
	//
	// There is no enum registry in Java. In v1, we used the enum registry
	// because enum types provided no reflective methods. The addition of
	// ProtoReflect removes that need.

	typesByName         typesByName
	extensionsByMessage extensionsByMessage

	numEnums      int
	numMessages   int
	numExtensions int
}

type (
	typesByName         map[protoreflect.FullName]Type
	extensionsByMessage map[protoreflect.FullName]extensionsByNumber
	extensionsByNumber  map[protoreflect.FieldNumber]protoreflect.ExtensionType
)

// NewTypes returns a registry initialized with the provided set of types.
// If there are conflicts, the first one takes precedence.
func NewTypes(typs ...Type) *Types {
	r := new(Types)
	r.Register(typs...) // ignore errors; first takes precedence
	return r
}

// Register registers the provided list of descriptor types.
//
// If a registration conflict occurs for enum, message, or extension types
// (e.g., two different types have the same full name),
// then the first type takes precedence and an error is returned.
func (r *Types) Register(typs ...Type) error {
	var firstErr error
typeLoop:
	for _, typ := range typs {
		switch typ.(type) {
		case protoreflect.EnumType, protoreflect.MessageType, protoreflect.ExtensionType:
			// Check for conflicts in typesByName.
			var desc protoreflect.Descriptor
			var pcnt *int
			switch t := typ.(type) {
			case protoreflect.EnumType:
				desc = t.Descriptor()
				pcnt = &r.numEnums
			case protoreflect.MessageType:
				desc = t.Descriptor()
				pcnt = &r.numMessages
			case protoreflect.ExtensionType:
				desc = t.TypeDescriptor()
				pcnt = &r.numExtensions
			default:
				panic(fmt.Sprintf("invalid type: %T", t))
			}
			name := desc.FullName()
			if prev := r.typesByName[name]; prev != nil {
				err := errors.New("%v %v is already registered", typeName(typ), name)
				err = amendErrorWithCaller(err, prev, typ)
				if r == GlobalTypes && ignoreConflict(desc, err) {
					err = nil
				}
				if firstErr == nil {
					firstErr = err
				}
				continue typeLoop
			}

			// Check for conflicts in extensionsByMessage.
			if xt, _ := typ.(protoreflect.ExtensionType); xt != nil {
				xd := xt.TypeDescriptor()
				field := xd.Number()
				message := xd.ContainingMessage().FullName()
				if prev := r.extensionsByMessage[message][field]; prev != nil {
					err := errors.New("extension number %d is already registered on message %v", field, message)
					err = amendErrorWithCaller(err, prev, typ)
					if r == GlobalTypes && ignoreConflict(xd, err) {
						err = nil
					}
					if firstErr == nil {
						firstErr = err
					}
					continue typeLoop
				}

				// Update extensionsByMessage.
				if r.extensionsByMessage == nil {
					r.extensionsByMessage = make(extensionsByMessage)
				}
				if r.extensionsByMessage[message] == nil {
					r.extensionsByMessage[message] = make(extensionsByNumber)
				}
				r.extensionsByMessage[message][field] = xt
			}

			// Update typesByName and the count.
			if r.typesByName == nil {
				r.typesByName = make(typesByName)
			}
			r.typesByName[name] = typ
			(*pcnt)++
		default:
			if firstErr == nil {
				firstErr = errors.New("invalid type: %v", typeName(typ))
			}
		}
	}
	return firstErr
}

// FindEnumByName looks up an enum by its full name.
// E.g., "google.protobuf.Field.Kind".
//
// This returns (nil, NotFound) if not found.
func (r *Types) FindEnumByName(enum protoreflect.FullName) (protoreflect.EnumType, error) {
	if r == nil {
		return nil, NotFound
	}
	if v := r.typesByName[enum]; v != nil {
		if et, _ := v.(protoreflect.EnumType); et != nil {
			return et, nil
		}
		return nil, errors.New("found wrong type: got %v, want enum", typeName(v))
	}
	return nil, NotFound
}

// FindMessageByName looks up a message by its full name.
// E.g., "google.protobuf.Any"
//
// This return (nil, NotFound) if not found.
func (r *Types) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
	// The full name by itself is a valid URL.
	return r.FindMessageByURL(string(message))
}

// FindMessageByURL looks up a message by a URL identifier.
// See documentation on google.protobuf.Any.type_url for the URL format.
//
// This returns (nil, NotFound) if not found.
func (r *Types) FindMessageByURL(url string) (protoreflect.MessageType, error) {
	if r == nil {
		return nil, NotFound
	}
	message := protoreflect.FullName(url)
	if i := strings.LastIndexByte(url, '/'); i >= 0 {
		message = message[i+len("/"):]
	}

	if v := r.typesByName[message]; v != nil {
		if mt, _ := v.(protoreflect.MessageType); mt != nil {
			return mt, nil
		}
		return nil, errors.New("found wrong type: got %v, want message", typeName(v))
	}
	return nil, NotFound
}

// FindExtensionByName looks up a extension field by the field's full name.
// Note that this is the full name of the field as determined by
// where the extension is declared and is unrelated to the full name of the
// message being extended.
//
// This returns (nil, NotFound) if not found.
func (r *Types) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
	if r == nil {
		return nil, NotFound
	}
	if v := r.typesByName[field]; v != nil {
		if xt, _ := v.(protoreflect.ExtensionType); xt != nil {
			return xt, nil
		}
		return nil, errors.New("found wrong type: got %v, want extension", typeName(v))
	}
	return nil, NotFound
}

// FindExtensionByNumber looks up a extension field by the field number
// within some parent message, identified by full name.
//
// This returns (nil, NotFound) if not found.
func (r *Types) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	if r == nil {
		return nil, NotFound
	}
	if xt, ok := r.extensionsByMessage[message][field]; ok {
		return xt, nil
	}
	return nil, NotFound
}

// NumEnums reports the number of registered enums.
func (r *Types) NumEnums() int {
	if r == nil {
		return 0
	}
	return r.numEnums
}

// RangeEnums iterates over all registered enums.
// Iteration order is undefined.
func (r *Types) RangeEnums(f func(protoreflect.EnumType) bool) {
	if r == nil {
		return
	}
	for _, typ := range r.typesByName {
		if et, ok := typ.(protoreflect.EnumType); ok {
			if !f(et) {
				return
			}
		}
	}
}

// NumMessages reports the number of registered messages.
func (r *Types) NumMessages() int {
	if r == nil {
		return 0
	}
	return r.numMessages
}

// RangeMessages iterates over all registered messages.
// Iteration order is undefined.
func (r *Types) RangeMessages(f func(protoreflect.MessageType) bool) {
	if r == nil {
		return
	}
	for _, typ := range r.typesByName {
		if mt, ok := typ.(protoreflect.MessageType); ok {
			if !f(mt) {
				return
			}
		}
	}
}

// NumExtensions reports the number of registered extensions.
func (r *Types) NumExtensions() int {
	if r == nil {
		return 0
	}
	return r.numExtensions
}

// RangeExtensions iterates over all registered extensions.
// Iteration order is undefined.
func (r *Types) RangeExtensions(f func(protoreflect.ExtensionType) bool) {
	if r == nil {
		return
	}
	for _, typ := range r.typesByName {
		if xt, ok := typ.(protoreflect.ExtensionType); ok {
			if !f(xt) {
				return
			}
		}
	}
}

// NumExtensionsByMessage reports the number of registered extensions for
// a given message type.
func (r *Types) NumExtensionsByMessage(message protoreflect.FullName) int {
	if r == nil {
		return 0
	}
	return len(r.extensionsByMessage[message])
}

// RangeExtensionsByMessage iterates over all registered extensions filtered
// by a given message type. Iteration order is undefined.
func (r *Types) RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool) {
	if r == nil {
		return
	}
	for _, xt := range r.extensionsByMessage[message] {
		if !f(xt) {
			return
		}
	}
}

func typeName(t Type) string {
	switch t.(type) {
	case protoreflect.EnumType:
		return "enum"
	case protoreflect.MessageType:
		return "message"
	case protoreflect.ExtensionType:
		return "extension"
	default:
		return fmt.Sprintf("%T", t)
	}
}

func amendErrorWithCaller(err error, prev, curr interface{}) error {
	prevPkg := goPackage(prev)
	currPkg := goPackage(curr)
	if prevPkg == "" || currPkg == "" || prevPkg == currPkg {
		return err
	}
	return errors.New("%s\n\tpreviously from: %q\n\tcurrently from:  %q", err, prevPkg, currPkg)
}

func goPackage(v interface{}) string {
	switch d := v.(type) {
	case protoreflect.EnumType:
		v = d.Descriptor()
	case protoreflect.MessageType:
		v = d.Descriptor()
	case protoreflect.ExtensionType:
		v = d.TypeDescriptor()
	}
	if d, ok := v.(protoreflect.Descriptor); ok {
		v = d.ParentFile()
	}
	if d, ok := v.(interface{ GoPackagePath() string }); ok {
		return d.GoPackagePath()
	}
	return ""
}
