package iterator

type Iterator interface {
	Current() string
	Next()
	HasNext() bool
}
