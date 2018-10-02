// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pretty

import (
	"testing"
)

func BenchmarkFromat(b *testing.B) {
	s := []byte(`{"string": "a", "number": 3, "array": [1, 2, 3], "map": {"map": "value"}, "emptyArray": [], "emptyMap": {}}`)
	f := NewFormatter()

	if _, err := f.Format(s); err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		f.Format(s)
	}
}
