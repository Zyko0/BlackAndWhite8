package core

import (
	"github.com/Zyko0/BlackAndWhite8/core/entity"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

const (
	MinAoeSpawnInterval = logic.TPS * 0.1 * 5
)

func (c *Core) handleAoeSpawn() {
	const center = logic.ScreenHeight / 2

	if c.ticks%MinAoeSpawnInterval == 0 {
		var x, y float32

		v := 0.1 + c.rng.Float32()*(logic.ScreenWidth-entity.AoeDefaultWidth-0.1)
		if c.rng.Intn(2) == 0 {
			if c.Player.X < center {
				x = logic.ScreenHeight
			}
			y = v
		} else {
			if c.Player.Y < center {
				y = logic.ScreenHeight
			}
			x = v
		}
		c.Aoes = append(c.Aoes, entity.NewAoe(
			x, y, entity.AoeDefaultWidth,
		))
	}
}

func (c *Core) handleProjectilesSpawn() {
	
}
