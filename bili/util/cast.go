package util

func ToPrt[T any](i T) *T {
	return &i
}

func ToVl[T any](i *T) T {
	return *i
}
