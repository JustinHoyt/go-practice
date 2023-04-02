package tap

import "fmt"

type Fn[T any] func(T) (T, error)

func Tap[T any](fn Fn[T]) Fn[T] {
	return func(x T) (T, error) {
		fn(x)
		return x, nil
	}
}

func UnaryPrintln[T any](val T) (int, error) {
	return fmt.Println(val)
}
