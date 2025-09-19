//go:build locafero_disable_conc

package pool

import "sync"

type Serial[T any] struct {
	results []T
	errors  error

	mu sync.Mutex
}

func New[T any]() *Serial[T] {
	return &Serial[T]{}
}

func (p *Serial[T]) Go(fn func() (T, error)) {
	p.mu.Lock()
	defer p.mu.Unlock()

	result, err := fn()
	if err != nil {
		p.errors = err
		return
	}

	p.results = append(p.results, result)
}

func (p *Serial[T]) Wait() ([]T, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.errors != nil {
		return nil, p.errors
	}

	results := p.results
	p.results = nil

	return results, nil
}
