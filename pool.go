package pool

import (
	"sync"
)

type pool[T any] struct {
	sync.Mutex
	ts   []T
	i, n int
	f    func() T
}

func New[T any](n int, f func() T) *pool[T] {
	if n == 0 {
		return nil
	}
	return &pool[T]{
		sync.Mutex{},
		make([]T, n),
		0,
		n,
		f,
	}
}
func (p *pool[T]) Get() (r T) {
	p.Lock()
	if p.i > 0 {
		p.i--
		r = p.ts[p.i]
		p.Unlock()
		return
	}
	p.Unlock()
	r = p.f()
	return
}
func (p *pool[T]) Put(r T) {
	p.Lock()
	if p.i < p.n {
		p.ts[p.i] = r
		p.i++
	}
	p.Unlock()
}
