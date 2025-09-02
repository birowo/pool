package pool

import (
	"sync"
)

type pool[T any] struct {
	sync.Mutex
	ts   []T
	i, n uint
	f    func() T
}

func New[T any](n uint, f func() T) *pool[T] {
	n = 1 << n
	if n == 0 {
		return nil
	}
	return &pool[T]{
		sync.Mutex{},
		make([]T, n),
		0, n - 1,
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
	p.i = (p.i + 1) & p.n
	p.Unlock()
}
