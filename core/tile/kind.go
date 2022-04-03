package tile

const (
	MethodEvenOdd int = iota
	MethodMin
	MethodMax
	MethodCount
)

type Kind struct {
	Method int
	Arg    int
}

var (
	kinds = []Kind{
		{MethodMin, 3},
		{MethodEvenOdd, 2},
		{MethodMin, 4},
		{MethodEvenOdd, 3},
		{MethodMin, 5},
		{MethodMin, 7},
		{MethodMax, 3},
		{MethodMax, 5},
	}
)

func MaxKind() int {
	return len(kinds) - 1
}

func GetKind(index int) Kind {
	return kinds[index]
}
