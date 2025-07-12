// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"internal/race"
	"sync/atomic"
	"unsafe"
)

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
//
// In the terminology of the Go memory model, a call to Done
// “synchronizes before” the return of any Wait call that it unblocks.
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.
//
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// If a WaitGroup is reused to wait for several independent sets of events,
// new Add calls must happen after all previous Wait calls have returned.
// See the WaitGroup example.
func (wg *WaitGroup) Add(delta int) {

	state := wg.state.Add(uint64(delta) << 32)
	v := int32(state >> 32)
	w := uint32(state)

	if v < 0 {
		panic("sync: negative WaitGroup counter")
	} //counter小于0了
	if w != 0 && delta > 0 && v == int32(delta) {
		//w != 0表示有协程在等待，而v == int32(delta)表示当前协程之前counter是0，说明ADD函数和wait函数存在并发！
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	if v > 0 || w == 0 {
		return
	}
	// This goroutine has set counter to 0 when waiters > 0.
	// Now there can't be concurrent mutations of state:
	// - Adds must not happen concurrently with Wait,
	// - Wait does not increment waiters if it sees counter == 0.
	// Still do a cheap sanity check to detect WaitGroup misuse.
	// 到这里说明该唤醒睡眠的协程了，因为这里一定counter==0且waiter>0
	// add函数并发不可能导致这个panic
	// add和wait存在并发，可能会导致panic，这里再简单的校验一下
	if wg.state.Load() != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// Reset waiters count to 0.
	wg.state.Store(0)
	for ; w != 0; w-- { //一个一个唤醒睡眠的协程
		runtime_Semrelease(&wg.sema, false, 0) //三个参数：1.信号量（就是个int地址）；2.锁调度策略开关，是否将执行权讲给唤醒的协程，waitGroup中不涉及；3.debug时候，和追踪挂钩
	}
}

// Wait blocks until the WaitGroup counter is zero.
func (wg *WaitGroup) Wait() {

	for {
		state := wg.state.Load()
		v := int32(state >> 32)
		w := uint32(state)
		if v == 0 {
			// Counter is 0, no need to wait.
			return
		}
		// Increment waiters count.
		if wg.state.CompareAndSwap(state, state+1) {

			runtime_Semacquire(&wg.sema)
			if wg.state.Load() != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}

			return
		}
	}
}
