package math

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}

	return x
}
