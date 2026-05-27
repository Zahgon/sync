// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package singleflight provides a duplicate function call suppression
// mechanism.
package singleflight // import "golang.org/x/sync/singleflight"

import (
	"errors"
	"sync"
)

// errGoexit indicates the runtime.Goexit was called in
// the user given function.
var errGoexit = errors.New("runtime.Goexit was called")

// A panicError is an arbitrary value recovered from a panic
// with the stack trace during the execution of given function.
type panicError struct {
	value any
	stack []byte
}

// Error implements error interface.
func (p *panicError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *panicError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newPanicError(v any) error { _ = "STUB: not implemented"; return nil }

// The first line of the stack trace is of the form "goroutine N [status]:"
// but by the time the panic reaches Do the goroutine may no longer exist
// and its status will have changed. Trim out the misleading line.

// call is an in-flight or completed singleflight.Do call
type call struct {
	wg sync.WaitGroup

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val any
	err error

	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups  int
	chans []chan<- Result
}

// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Result holds the results of Do, so they can be passed
// on a channel.
type Result struct {
	Val    any
	Err    error
	Shared bool
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
func (g *Group) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
	_ = "STUB: not implemented"
	return *new(any), nil, false
}

// DoChan is like Do but returns a channel that will receive the
// results when they are ready.
//
// The returned channel will not be closed.
func (g *Group) DoChan(key string, fn func() (any, error)) <-chan Result {
	_ = "STUB: not implemented"
	return nil
}

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
	_ = "STUB: not implemented"
	return
}

// use double-defer to distinguish panic from runtime.Goexit,
// more details see https://golang.org/cl/134395

// the given function invoked runtime.Goexit

// In order to prevent the waiting channels from being blocked forever,
// needs to ensure that this panic cannot be recovered.

// Keep this goroutine around so that it will appear in the crash dump.

// Already in the process of goexit, no need to call again

// Normal return

// Ideally, we would wait to take a stack trace until we've determined
// whether this is a panic or a runtime.Goexit.
//
// Unfortunately, the only way we can distinguish the two is to see
// whether the recover stopped the goroutine from terminating, and by
// the time we know that, the part of the stack trace relevant to the
// panic has been discarded.

// Forget tells the singleflight to forget about a key.  Future calls
// to Do for this key will call the function rather than waiting for
// an earlier call to complete.
func (g *Group) Forget(key string) { _ = "STUB: not implemented"; return }
