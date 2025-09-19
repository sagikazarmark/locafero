package queue

type Queue[T any] interface {
	Add(func() (T, error))
	Wait() ([]T, error)
}
