package other

type One[T any] struct{}

type Two[T any, R any] struct{}

type Three struct{}

type Four struct{}

type Five interface{}

type Twenty[T any, R any] interface {
	Twenty() (T, R)
}

type TwentyThree[T any] interface {
	TwentyThree() StructType
}

type StructType struct{}
