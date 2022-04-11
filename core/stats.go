package core

import "time"

type Statistics struct {
	LoopCount  int
	Duration   time.Duration
	Completion float64
	HP         int
}

func (c *Core) GetStatistics() Statistics {
	return Statistics{
		LoopCount:  c.loop,
		Duration:   c.GetTime(),
		Completion: c.GetCompletion(),
		HP:         c.Player.HP,
	}
}
