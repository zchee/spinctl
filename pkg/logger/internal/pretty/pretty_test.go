// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pretty

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func TestMarshal(t *testing.T) {
	prettyJSON := func(s string) string {
		var v interface{}

		err := json.Unmarshal([]byte(s), &v)

		if err != nil {
			t.Error(err)
		}

		prettyJSONByte, err := Marshal(v)

		if err != nil {
			t.Error(err)
		}

		return string(prettyJSONByte)
	}

	test := func(expected, actual string) {
		if expected != actual {
			t.Errorf("\nexpected:\n%s\n\nactual:\n%s", expected, actual)
		}
	}

	blueBold := color.New(color.FgBlue, color.Bold).SprintFunc()
	greenBold := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	blackBold := color.New(color.FgBlack, color.Bold).SprintFunc()
	yelloBold := color.New(color.FgYellow, color.Bold).SprintFunc()

	actual := prettyJSON(`{
  "key": {
    "a": "str",
    "b": 100,
    "c": null,
    "d": true,
    "e": false,
    "f": { "key": "str" },
	"g": {},
	"h": []
  }
}`)

	expectedFormat := `{
  %s: {
    %s: %s,
    %s: %s,
    %s: %s,
    %s: %s,
    %s: %s,
    %s: {
      %s: %s
    },
    %s: {},
    %s: []
  }
}`

	expected := fmt.Sprintf(expectedFormat,
		blueBold(`"key"`),
		blueBold(`"a"`), greenBold(`"str"`),
		blueBold(`"b"`), cyanBold("100"),
		blueBold(`"c"`), blackBold("null"),
		blueBold(`"d"`), yelloBold("true"),
		blueBold(`"e"`), yelloBold("false"),
		blueBold(`"f"`), blueBold(`"key"`), greenBold(`"str"`),
		blueBold(`"g"`),
		blueBold(`"h"`),
	)

	test(expected, actual)
	test("{}", prettyJSON("{}"))
	test("[]", prettyJSON("[]"))
}

func TestStringEscape(t *testing.T) {
	f := NewFormatter()
	f.DisabledColor = true
	s := `{"foo":"foo\"\nbar"}`
	r, err := f.Format([]byte(s))

	if err != nil {
		t.Error(err)
	}

	expected := `{
  "foo": "foo\"\nbar"
}`

	if string(r) != expected {
		t.Errorf("actual: %s\nexpected: %s", string(r), expected)
	}
}

func TestStringPercentEscape(t *testing.T) {
	f := NewFormatter()
	s := `{"foo":"foo%2Fbar"}`
	r, err := f.Format([]byte(s))

	if err != nil {
		t.Error(err)
	}

	expectedFormat := `{
  %s: %s
}`

	blueBold := color.New(color.FgBlue, color.Bold).SprintFunc()
	greenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

	expected := fmt.Sprintf(expectedFormat,
		blueBold(`"foo"`), greenBold(`"foo%2Fbar"`),
	)

	if string(r) != expected {
		t.Errorf("actual: %s\nexpected: %s", string(r), expected)
	}
}

func TestStringPercentEscape_DisabledColor(t *testing.T) {
	f := NewFormatter()
	f.DisabledColor = true
	s := `{"foo":"foo%2Fbar"}`
	r, err := f.Format([]byte(s))

	if err != nil {
		t.Error(err)
	}

	expected := `{
  "foo": "foo%2Fbar"
}`

	if string(r) != expected {
		t.Errorf("actual: %s\nexpected: %s", string(r), expected)
	}
}
