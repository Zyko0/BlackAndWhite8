package core

import (
	"image"
	"math/rand"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/assets/shape"
	"github.com/Zyko0/BlackAndWhite8/core/entity"
	"github.com/Zyko0/BlackAndWhite8/core/utils"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

const (
	ticksMaxDifficulty = 3 * logic.TPS * 60 // Minutes
)

type Core struct {
	ticks              uint64
	rng                *rand.Rand
	loop               int
	aoeInterval        uint64
	projectileInterval uint64

	clock      time.Time
	paused     bool
	pauseClock time.Time

	Shape       *shape.Shape
	Board       *Board
	Player      *Player
	Aoes        []*entity.Aoe
	Projectiles []*entity.Projectile
}

var autoed = false // TODO: remove

func New() *Core {
	autoed = false // TODO: remove
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	s := shape.Random(rng)
	s.ApplyRandomRotation(rng)

	now := time.Now()

	return &Core{
		rng:                rng,
		aoeInterval:        initialAoeSpawnInterval,
		projectileInterval: initialProjectileSpawnInterval,

		clock:      now,
		pauseClock: now,

		Shape:  s,
		Board:  newBoard(rng, s),
		Player: newPlayer(),
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
		assets.PlayDashSound()
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
	// Ignore previous intents if players is being knocked back
	if c.Player.KnockbackDuration > 0 {
		dx, dy = c.Player.knockbackDx, c.Player.knockbackDy
		dx *= KnockbackSpeed
		dy *= KnockbackSpeed
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
			tile.Completed = (tile.KindIndex == c.Shape.At(int(tile.X), int(tile.Y)))
			if tile.Completed {
				assets.PlayFlipSound()
			} else {
				assets.PlayFlipFailSound()
			}
		}
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

				dx, dy := utils.GetKnockbackVector(playerRect, rect)
				c.Player.knockbackDx = dx
				c.Player.knockbackDy = dy
				c.Player.TakeDamage()
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

				dx, dy := utils.GetKnockbackVector(playerRect, rect)
				c.Player.knockbackDx = dx
				c.Player.knockbackDy = dy
				c.Player.TakeDamage()
				break
			}
		}
	}
}

func (c *Core) Loop() {
	c.loop++
	ticks := uint64(c.loop) * (ticksMaxDifficulty / 10)
	// Note: if loop is spammed, do not increase difficulty, or should ?
	if c.ticks > ticks {
		c.ticks = ticks
	}

	t := c.Board.TileAt(c.Player.X, c.Player.Y)
	c.Player.X = float32(t.X)*t.W + t.W/2 - PlayerSize/2
	c.Player.Y = float32(t.Y)*t.H + t.H/2 - PlayerSize/2

	c.Aoes = c.Aoes[:0]
	c.Projectiles = c.Projectiles[:0]
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
	// Adjust spawning rates
	ratio := float64(c.ticks) / ticksMaxDifficulty
	if ratio > 1 {
		ratio = 1
	}
	c.aoeInterval = initialAoeSpawnInterval - uint64(float64(initialAoeSpawnInterval-minAoeSpawnInterval)*ratio)
	c.projectileInterval = initialProjectileSpawnInterval - uint64(float64(initialProjectileSpawnInterval-minProjectileSpawnInterval)*ratio)
	// Player
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

func (c *Core) TogglePause() {
	c.paused = !c.paused
	if c.paused {
		// Pause
		c.pauseClock = time.Now()
		assets.StopMusic()
		assets.ResumeMenuMusic()
	} else {
		// Resume
		c.clock = c.clock.Add(time.Since(c.pauseClock))
		assets.StopMenuMusic()
		assets.ResumeMusic()
	}
}

func (c *Core) IsPaused() bool {
	return c.paused
}

func (c *Core) GetTime() time.Duration {
	if !c.paused {
		c.pauseClock = time.Now()
	}
	return c.pauseClock.Sub(c.clock)
}

func (c *Core) GetLoop() int {
	return c.loop
}

func (c *Core) GetCompletion() float64 {
	return float64(c.Board.completed) / float64(c.Board.Size*c.Board.Size)
}

func (c *Core) GetDifficulty() float64 {
	return float64(c.ticks) / ticksMaxDifficulty
}

func (c *Core) IsOver() bool {
	return c.Player.HP == 0 || c.Board.completed == c.Board.Size*c.Board.Size
}
