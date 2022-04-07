package shape

import (
	"math/rand"

	"github.com/Zyko0/BlackAndWhite8/core/tile"
)

func (s *Shape) Size() int {
	return len(s.values)
}

func (s *Shape) At(x, y int) int {
	return s.values[y][x]
}

func (s *Shape) ApplyRandomRotation(rng *rand.Rand) {
	maxKind := tile.MaxKind()
	rotation := rng.Intn(maxKind) + 1
	for y, row := range s.values {
		for x, v := range row {
			nv := v + rotation
			if nv > maxKind {
				nv -= maxKind
			}
			s.values[y][x] = nv
		}
	}
}

const (
	SizeX16 = 16
	SizeX32 = 32
)

func Random(rng *rand.Rand, size int) *Shape {
	var s Shape

	switch size {
	case SizeX16:
		s = shapes16[rng.Intn(len(shapes16))]
	case SizeX32:
		s = shapes32[rng.Intn(len(shapes32))]
	default:
		panic("invalid size")
	}

	values := make([][]int, len(s.values))
	for i := range s.values {
		values[i] = make([]int, len(s.values[i]))
		copy(values[i], s.values[i])
	}

	return &Shape{
		maxValue: s.maxValue,
		minValue: s.minValue,
		values:   values,
	}
}
