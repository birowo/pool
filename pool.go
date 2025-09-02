package pool

import (
	"sync"
)

type pool[T any] struct {
	sync.Mutex
	ts []T
	i  int
	f  func() T
}

func New[T any](n int, f func() T) *pool[T] {
	return &pool[T]{
		sync.Mutex{},
		make([]T, 1<<n),
		0,
		f,
	}
}
func (p *pool[T]) Get() (r T) {
	p.Lock()
	if p.i == 0 {
		r = p.f()
	} else {
		p.i--
		r = p.ts[p.i]
	}
	p.Unlock()
	return
}
func (p *pool[T]) Put(r T) {
	p.Lock()
	p.ts[p.i] = r
	p.i++
	p.Unlock()
}
