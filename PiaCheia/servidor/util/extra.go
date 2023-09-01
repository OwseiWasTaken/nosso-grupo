package util

func Panic(e error) {
	if (e != nil) {
		panic(e)
	}
}

func Unpack[T any](v T, e error) T {
	if (e != nil) {
		panic(e)
	}
	return v
}
