// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"bytes"
	"io"
	"regexp"
	"strconv"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type syntaxError struct{ error }

func newSyntaxError(f string, x ...interface{}) error {
	return syntaxError{errors.New(f, x...)}
}

// Unmarshal parses b as the proto text format.
// It returns a Value, which is always of the Message type.
func Unmarshal(b []byte) (Value, error) {
	p := decoder{in: b}
	p.consume(0) // trim leading spaces or comments
	v, err := p.unmarshalMessage(false)
	if err != nil {
		if e, ok := err.(syntaxError); ok {
			b = b[:len(b)-len(p.in)] // consumed input
			line := bytes.Count(b, []byte("\n")) + 1
			if i := bytes.LastIndexByte(b, '\n'); i >= 0 {
				b = b[i+1:]
			}
			column := utf8.RuneCount(b) + 1 // ignore multi-rune characters
			err = errors.New("syntax error (line %d:%d): %v", line, column, e.error)
		}
		return Value{}, err
	}
	if len(p.in) > 0 {
		return Value{}, errors.New("%d bytes of unconsumed input", len(p.in))
	}
	return v, nil
}

type decoder struct {
	in []byte
}

func (p *decoder) unmarshalList() (Value, error) {
	b := p.in
	var elems []Value
	if err := p.consumeChar('[', "at start of list"); err != nil {
		return Value{}, err
	}
	if len(p.in) > 0 && p.in[0] != ']' {
		for len(p.in) > 0 {
			v, err := p.unmarshalValue()
			if err != nil {
				return Value{}, err
			}
			elems = append(elems, v)
			if !p.tryConsumeChar(',') {
				break
			}
		}
	}
	if err := p.consumeChar(']', "at end of list"); err != nil {
		return Value{}, err
	}
	b = b[:len(b)-len(p.in)]
	return rawValueOf(elems, b[:len(b):len(b)]), nil
}

func (p *decoder) unmarshalMessage(checkDelims bool) (Value, error) {
	b := p.in
	var items [][2]Value
	delims := [2]byte{'{', '}'}
	if len(p.in) > 0 && p.in[0] == '<' {
		delims = [2]byte{'<', '>'}
	}
	if checkDelims {
		if err := p.consumeChar(delims[0], "at start of message"); err != nil {
			return Value{}, err
		}
	}
	for len(p.in) > 0 {
		if p.in[0] == '}' || p.in[0] == '>' {
			break
		}
		k, err := p.unmarshalKey()
		if err != nil {
			return Value{}, err
		}
		if !p.tryConsumeChar(':') && len(p.in) > 0 && p.in[0] != '{' && p.in[0] != '<' {
			return Value{}, newSyntaxError("expected ':' after message key")
		}
		v, err := p.unmarshalValue()
		if err != nil {
			return Value{}, err
		}
		if p.tryConsumeChar(';') || p.tryConsumeChar(',') {
			// always optional
		}
		items = append(items, [2]Value{k, v})
	}
	if checkDelims {
		if err := p.consumeChar(delims[1], "at end of message"); err != nil {
			return Value{}, err
		}
	}
	b = b[:len(b)-len(p.in)]
	return rawValueOf(items, b[:len(b):len(b)]), nil
}

// unmarshalKey parses the key, which may be a Name, String, or Uint.
func (p *decoder) unmarshalKey() (v Value, err error) {
	if p.tryConsumeChar('[') {
		if len(p.in) == 0 {
			return Value{}, io.ErrUnexpectedEOF
		}
		if p.in[0] == '\'' || p.in[0] == '"' {
			// Historically, Go's parser allowed a string for the Any type URL.
			// This is specific to Go and contrary to the C++ implementation,
			// which does not support strings for the Any type URL.
			v, err = p.unmarshalString()
			if err != nil {
				return Value{}, err
			}
		} else {
			v, err = p.unmarshalURL()
			if err != nil {
				return Value{}, err
			}
		}
		if err := p.consumeChar(']', "at end of extension name"); err != nil {
			return Value{}, err
		}
		return v, nil
	}
	v, err = p.unmarshalName()
	if err == nil {
		return v, nil
	}
	v, err = p.unmarshalNumberKey()
	if err == nil {
		return v, nil
	}
	return Value{}, err
}

// unmarshalURL parses an Any type URL string. The C++ parser does not handle
// many legal URL strings. This implementation is more liberal and allows for
// the pattern ^[-_a-zA-Z0-9]+([./][-_a-zA-Z0-9]+)*`).
func (p *decoder) unmarshalURL() (Value, error) {
	s := p.in
	var size int
	for len(s) > 0 && (s[0] == '-' || s[0] == '_' ||
		('0' <= s[0] && s[0] <= '9') ||
		('a' <= s[0] && s[0] <= 'z') ||
		('A' <= s[0] && s[0] <= 'Z')) {
		s = s[1:]
		size++
		if len(s) > 0 && (s[0] == '/' || s[0] == '.') {
			s = s[1:]
			size++
		}
	}

	// Last character cannot be '.' or '/'.
	// Next byte should either be a delimiter or it is at the end.
	if size == 0 || p.in[size-1] == '.' || p.in[size-1] == '/' ||
		(len(s) > 0 && !isDelim(s[0])) {
		return Value{}, newSyntaxError("invalid %q as identifier", errRegexp.Find(p.in))
	}
	v := rawValueOf(string(p.in[:size]), p.in[:size:size])
	p.consume(size)
	return v, nil
}

// unmarshalNumberKey parses field number as key. Field numbers are non-negative
// integers.
func (p *decoder) unmarshalNumberKey() (Value, error) {
	num, ok := parseNumber(p.in)
	if !ok || num.neg || num.typ == numFloat {
		return Value{}, newSyntaxError("invalid %q as identifier", errRegexp.Find(p.in))
	}
	v, err := strconv.ParseUint(string(num.value), 0, 64)
	if err != nil {
		return Value{}, newSyntaxError("invalid %q as identifier", errRegexp.Find(p.in))
	}
	p.consume(num.size)
	return rawValueOf(v, num.value), nil
}

func (p *decoder) unmarshalValue() (Value, error) {
	if len(p.in) == 0 {
		return Value{}, io.ErrUnexpectedEOF
	}
	switch p.in[0] {
	case '"', '\'':
		return p.unmarshalStrings()
	case '[':
		return p.unmarshalList()
	case '{', '<':
		return p.unmarshalMessage(true)
	default:
		n, ok := consumeName(p.in)
		if ok && literals[string(p.in[:n])] == nil {
			v := rawValueOf(protoreflect.Name(p.in[:n]), p.in[:n:n])
			p.consume(n)
			return v, nil
		}
		return p.unmarshalNumber()
	}
}

// unmarshalName unmarshals an unquoted proto identifier.
// Regular expression that matches an identifier: `^[_a-zA-Z][_a-zA-Z0-9]*`
//
// E.g., `field_name` => ValueOf(protoreflect.Name("field_name"))
func (p *decoder) unmarshalName() (Value, error) {
	n, ok := consumeName(p.in)
	if !ok {
		return Value{}, newSyntaxError("invalid %q as identifier", errRegexp.Find(p.in))
	}

	v := rawValueOf(protoreflect.Name(p.in[:n]), p.in[:n:n])
	p.consume(n)
	return v, nil
}

func consumeName(input []byte) (int, bool) {
	var n int

	s := input
	if len(s) == 0 {
		return 0, false
	}

	switch {
	case s[0] == '_',
		'a' <= s[0] && s[0] <= 'z',
		'A' <= s[0] && s[0] <= 'Z':
		s = s[1:]
		n++
	default:
		return 0, false
	}

	for len(s) > 0 && (s[0] == '_' ||
		'a' <= s[0] && s[0] <= 'z' ||
		'A' <= s[0] && s[0] <= 'Z' ||
		'0' <= s[0] && s[0] <= '9') {
		s = s[1:]
		n++
	}

	if len(s) > 0 && !isDelim(s[0]) {
		return 0, false
	}

	return n, true
}

func (p *decoder) consumeChar(c byte, msg string) error {
	if p.tryConsumeChar(c) {
		return nil
	}
	if len(p.in) == 0 {
		return io.ErrUnexpectedEOF
	}
	return newSyntaxError("invalid character %q, expected %q %s", p.in[0], c, msg)
}

func (p *decoder) tryConsumeChar(c byte) bool {
	if len(p.in) > 0 && p.in[0] == c {
		p.consume(1)
		return true
	}
	return false
}

// consume consumes n bytes of input and any subsequent whitespace or comments.
func (p *decoder) consume(n int) {
	p.in = p.in[n:]
	for len(p.in) > 0 {
		switch p.in[0] {
		case ' ', '\n', '\r', '\t':
			p.in = p.in[1:]
		case '#':
			if i := bytes.IndexByte(p.in, '\n'); i >= 0 {
				p.in = p.in[i+len("\n"):]
			} else {
				p.in = nil
			}
		default:
			return
		}
	}
}

// Any sequence that looks like a non-delimiter (for error reporting).
var errRegexp = regexp.MustCompile(`^([-+._a-zA-Z0-9\/]+|.)`)

// isDelim returns true if given byte is a delimiter character.
func isDelim(c byte) bool {
	return !(c == '-' || c == '+' || c == '.' || c == '_' ||
		('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9'))
}
