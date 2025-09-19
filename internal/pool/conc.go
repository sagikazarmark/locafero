//go:build !locafero_disable_conc

package pool

import "github.com/sourcegraph/conc/pool"

type Pool[T any] = pool.ResultErrorPool[T]

func New[T any]() *Pool[T] {
	// Arbitrary go routine limit (TODO: make this a parameter)
	return pool.NewWithResults[T]().WithMaxGoroutines(5).WithErrors().WithFirstError()
}
