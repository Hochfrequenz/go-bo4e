package internal

func Ptr[S any](s S) *S {
	return &s
}
