// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"sync"
)

type SyncBuffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

func (s *SyncBuffer) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.buffer.Reset()
}

func (s *SyncBuffer) Write(p []byte) (n int, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.Write(p)
}

func (s *SyncBuffer) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.String()
}

func (s *SyncBuffer) Sync() error {
	return nil
}
