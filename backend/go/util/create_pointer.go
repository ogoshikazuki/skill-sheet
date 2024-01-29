package util

func CreatePointer[T any](input T) *T {
	return &input
}
