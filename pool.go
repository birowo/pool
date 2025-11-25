package pool

import (
	"sync/atomic"
)

type Pool[T any] struct {
	ts []T
	atomic.Int32
	int32
	fn func() T
}

func New[T any](n int32, fn func() T) *Pool[T] {
	if n == 0 {
		return nil
	}
	return &Pool[T]{
		make([]T, n),
		atomic.Int32{},
		n,
		fn,
	}
}
func (p *Pool[T]) Get() T {
	if p.CompareAndSwap(0, 0) {
		return p.fn()
	}
	return p.ts[p.Add(-1)]
}
func (p *Pool[T]) Put(x T) {
	if p.CompareAndSwap(p.int32, p.int32) {
		return
	}
	p.ts[p.Add(1)-1] = x
}
