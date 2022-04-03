package core

import (
	"math/rand"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets/shape"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

type Core struct {
	start     time.Time
	loopCount int

	Difficulty Difficulty
	Shape      *shape.Shape
	Board      *Board
	Player     *Player
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
		start: time.Now(),

		Difficulty: difficulty,
		Shape:      shape.Random(rng, size),
		Board:      newBoard(rng, size),
		Player:     newPlayer(),
	}
}

func (c *Core) handlePlayerCollisions() {
	dx, dy := c.Player.intentX*MoveSpeed, c.Player.intentY*MoveSpeed
	if c.Player.X+dx < 0 {
		dx -= c.Player.X
	}
	if c.Player.Y+dy < 0 {
		dy -= c.Player.Y
	}
	if c.Player.X+PlayerSize+dx > logic.ScreenHeight {
		dx -= (c.Player.X + PlayerSize) - logic.ScreenHeight
	}
	if c.Player.Y+dy+PlayerSize > logic.ScreenHeight {
		dy -= (c.Player.Y + PlayerSize) - logic.ScreenHeight
	}

	c.Player.X += dx
	c.Player.Y += dy
}

func (c *Core) Update() {
	// TODO: below code resolves the shape
	/*if !autoed {
		for y, row := range c.Board.Tiles {
			for x, tile := range row {
				tile.KindIndex = c.Shape.At(x, y)
			}
		}
		autoed = true
	}*/

	c.Player.Update()
	c.handlePlayerCollisions()

	c.Board.Update()
}

func (c *Core) GetTime() time.Duration {
	return time.Since(c.start)
}

func (c *Core) GetLoopCount() int {
	return c.loopCount
}

func (c *Core) GetProgression() float64 {
	// TODO: do
	return 0.
}
