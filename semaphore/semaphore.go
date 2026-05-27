// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semaphore provides a weighted semaphore implementation.
package semaphore // import "golang.org/x/sync/semaphore"

import (
	"container/list"
	"context"
	"sync"
)

type waiter struct {
	n     int64
	ready chan<- struct{} // Closed when semaphore acquired.
}

// NewWeighted creates a new weighted semaphore with the given
// maximum combined weight for concurrent access.
func NewWeighted(n int64) *Weighted { _ = "STUB: not implemented"; return nil }

// Weighted provides a way to bound concurrent access to a resource.
// The callers can request access with a given weight.
type Weighted struct {
	size    int64
	cur     int64
	mu      sync.Mutex
	waiters list.List
}

// Acquire acquires the semaphore with a weight of n, blocking until resources
// are available or ctx is done. On success, returns nil. On failure, returns
// ctx.Err() and leaves the semaphore unchanged.
func (s *Weighted) Acquire(ctx context.Context, n int64) error {
	_ = "STUB: not implemented"
	return nil
}

// ctx becoming done has "happened before" acquiring the semaphore,
// whether it became done before the call began or while we were
// waiting for the mutex. We prefer to fail even if we could acquire
// the mutex without blocking.

// Since we hold s.mu and haven't synchronized since checking done, if
// ctx becomes done before we return here, it becoming done must have
// "happened concurrently" with this call - it cannot "happen before"
// we return in this branch. So, we're ok to always acquire here.

// Don't make other Acquire calls block on one that's doomed to fail.

// Acquired the semaphore after we were canceled.
// Pretend we didn't and put the tokens back.

// If we're at the front and there're extra tokens left, notify other waiters.

// Acquired the semaphore. Check that ctx isn't already done.
// We check the done channel instead of calling ctx.Err because we
// already have the channel, and ctx.Err is O(n) with the nesting
// depth of ctx.

// TryAcquire acquires the semaphore with a weight of n without blocking.
// On success, returns true. On failure, returns false and leaves the semaphore unchanged.
func (s *Weighted) TryAcquire(n int64) bool { _ = "STUB: not implemented"; return false }

// Release releases the semaphore with a weight of n.
func (s *Weighted) Release(n int64) { _ = "STUB: not implemented"; return }

func (s *Weighted) notifyWaiters() { _ = "STUB: not implemented"; return }

// No more waiters blocked.

// Not enough tokens for the next waiter.  We could keep going (to try to
// find a waiter with a smaller request), but under load that could cause
// starvation for large requests; instead, we leave all remaining waiters
// blocked.
//
// Consider a semaphore used as a read-write lock, with N tokens, N
// readers, and one writer.  Each reader can Acquire(1) to obtain a read
// lock.  The writer can Acquire(N) to obtain a write lock, excluding all
// of the readers.  If we allow the readers to jump ahead in the queue,
// the writer will starve — there is always one token available for every
// reader.
