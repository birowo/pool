package pool

import (
	"sync/atomic"
)

type pool[T any] struct {
	ts []T
	atomic.Int32
	int32
	f func() T
}

func New[T any](n int32, f func() T) *pool[T] {
	n = 1 << n
	if n < 1 {
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
	i := p.Add(-1)
	if i < 0 {
		p.Store(0)
		r = p.f()
	} else {
		r = p.ts[i]
	}
	return
}
func (p *pool[T]) Put(r T) {
	i := p.Add(1) - 1
	if i > p.int32 {
		p.Store(p.int32)
		return
	}
	p.ts[i] = r
}
