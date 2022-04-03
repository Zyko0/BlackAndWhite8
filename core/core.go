package core

import (
	"math/rand"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets/shape"
)

type Core struct {
	Difficulty Difficulty
	Shape      *shape.Shape
	Board      *Board
}

var autoed = false // TODO: remove

func New(difficulty Difficulty) *Core {
	autoed = false // TODO: remove
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	size := shape.SizeX16
	if difficulty == DifficultyHard {
		size = shape.SizeX32
	}

	return &Core{
		Difficulty: difficulty,
		Shape:      shape.Random(rng, size),
		Board:      newBoard(rng, difficulty),
	}
}

func (c *Core) Update() {
	/*if !autoed {
		for y, row := range c.Board.Tiles {
			for x, tile := range row {
				tile.KindIndex = c.Shape.At(x, y)
			}
		}
		autoed = true
	}*/

	c.Board.Update()
}
