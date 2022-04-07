package core

import (
	"image"
	"math/rand"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets/shape"
	"github.com/Zyko0/BlackAndWhite8/core/entity"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

type Core struct {
	ticks     uint64
	rng       *rand.Rand
	start     time.Time
	loopCount int

	Difficulty  Difficulty
	Shape       *shape.Shape
	Board       *Board
	Player      *Player
	Aoes        []*entity.Aoe
	Projectiles []*entity.Projectile
}

var autoed = false // TODO: remove

func New(difficulty Difficulty) *Core {
	autoed = false // TODO: remove
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	size := shape.SizeX16
	if difficulty == DifficultyHard {
		size = shape.SizeX32
	}

	s := shape.Random(rng, size)
	s.ApplyRandomRotation(rng)

	return &Core{
		rng:   rng,
		start: time.Now(),

		Difficulty: difficulty,
		Shape:      s,
		Board:      newBoard(rng, s),
		Player:     newPlayer(),
	}
}

func (c *Core) handlePlayerIntents() {
	dx, dy := float32(c.Player.intentX), float32(c.Player.intentY)

	if c.Player.intentX != 0 && c.Player.intentY != 0 {
		dx *= 0.7071067
		dy *= 0.7071067
	}
	if c.Player.intentDash && c.Player.DashCD == 0 {
		c.Player.DashCD = DashCooldown
		c.Player.DashDuration = DashDuration
		// TODO: handle dash
	}
	if c.Player.DashCD > 0 {
		c.Player.DashCD--
	}
	if c.Player.DashDuration > 0 {
		dx *= DashSpeed
		dy *= DashSpeed
		c.Player.DashDuration--
	} else {
		dx *= MoveSpeed
		dy *= MoveSpeed
	}
	c.Player.X += dx
	c.Player.Y += dy

	if tile := c.Board.TileAt(c.Player.X+PlayerSize/2, c.Player.Y+PlayerSize/2); tile != nil {
		if c.Player.intentFlip {
			if tile.Highlighted {
				tile.KindIndex = c.Shape.At(int(tile.X), int(tile.Y))
			} else {
				tile.FlipUp()
			}
		}
		tile.Completed = (tile.KindIndex == c.Shape.At(int(tile.X), int(tile.Y)))
	}

	if c.Player.KnockbackDuration > 0 {
		c.Player.KnockbackDuration--
	}
	if c.Player.InvulnDuration > 0 {
		c.Player.InvulnDuration--
	}
}

func (c *Core) handlePlayerCollisions() {
	if c.Player.X < 0 {
		c.Player.X = 0
	}
	if c.Player.Y < 0 {
		c.Player.Y = 0
	}
	if c.Player.X+PlayerSize > logic.ScreenHeight {
		c.Player.X = logic.ScreenHeight - PlayerSize
	}
	if c.Player.Y+PlayerSize > logic.ScreenHeight {
		c.Player.Y = logic.ScreenHeight - PlayerSize
	}

	// Aoes check
	if c.Player.InvulnDuration == 0 {
		playerRect := image.Rect(int(c.Player.X), int(c.Player.Y), int(c.Player.X+PlayerSize), int(c.Player.Y+PlayerSize))
		for _, aoe := range c.Aoes {
			if aoe.IsOver() {
				continue
			}
			rect := aoe.GetRect()
			if rect.Overlaps(playerRect) {
				c.Player.InvulnDuration = InvulnTime
				c.Player.KnockbackDuration = KnockbackTime
				break
			}
		}
	}
	// Projectiles check
	if c.Player.InvulnDuration == 0 {
		playerRect := image.Rect(int(c.Player.X), int(c.Player.Y), int(c.Player.X+PlayerSize), int(c.Player.Y+PlayerSize))
		for _, proj := range c.Projectiles {
			rect := proj.GetRect()
			if rect.Overlaps(playerRect) {
				c.Player.InvulnDuration = InvulnTime
				c.Player.KnockbackDuration = KnockbackTime
				break
			}
		}
	}
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
	c.handlePlayerIntents()
	c.handlePlayerCollisions()
	// Entities
	c.handleAoeSpawn()
	for i := 0; i < len(c.Aoes); i++ {
		aoe := c.Aoes[i]
		if aoe.IsOver() {
			c.Aoes[i] = c.Aoes[len(c.Aoes)-1]
			c.Aoes = c.Aoes[:len(c.Aoes)-1]
		} else {
			aoe.Update()
		}
	}
	c.handleProjectilesSpawn()
	for i := 0; i < len(c.Projectiles); i++ {
		proj := c.Projectiles[i]
		x0, y0 := proj.X-entity.ProjectileRadius, proj.Y-entity.ProjectileRadius
		x1, y1 := proj.X+entity.ProjectileRadius, proj.Y+entity.ProjectileRadius
		if x1 < 0 || x0 > logic.ScreenHeight || y1 < 0 || y0 > logic.ScreenHeight {
			c.Projectiles[i] = c.Projectiles[len(c.Projectiles)-1]
			c.Projectiles = c.Projectiles[:len(c.Projectiles)-1]
		} else {
			proj.Update()
		}
	}

	c.Board.Update(c.Shape)

	c.ticks++
}

func (c *Core) GetTime() time.Duration {
	return time.Since(c.start)
}

func (c *Core) GetLoopCount() int {
	return c.loopCount
}

func (c *Core) GetCompletion() float64 {
	return float64(c.Board.completed) / float64(c.Board.Size*c.Board.Size)
}
