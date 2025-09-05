package pool

import (
	"sync/atomic"
)

type pool[T any] struct {
	ts []T
	atomic.Int32
	n int32
	f func() T
}

func New[T any](n int32, f func() T) *pool[T] {
	if n == 0 {
		return nil
	}
	return &pool[T]{
		make([]T, n),
		atomic.Int32{},
		n,
		f,
	}
}
func (p *pool[T]) Get() (r T) {
	if p.CompareAndSwap(0, 0) {
		r = p.f()
		return
	}
	r = p.ts[p.Add(-1)]
	return
}
func (p *pool[T]) Put(r T) (ret bool) {
	ret = p.CompareAndSwap(p.n, p.n)
	if ret {
		return
	}
	p.ts[p.Add(1)-1] = r
	return
}
