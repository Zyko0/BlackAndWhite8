package graphics

import (
	"math"
)

type Loop struct {
	ticks float64

	currentScale float64
	bgScale      float64
	tx, ty       float64

	Done bool
}

func (l *Loop) Update() {
	const (
		zoomIntensity = 0.025
	)

	l.currentScale = math.Exp(zoomIntensity * l.ticks)
	l.bgScale -= math.Abs(math.Sin(l.ticks*math.Pi/330)) * 0.3 // 0.1
	if l.bgScale < 1 {
		l.bgScale = 1
		l.Done = true
	}

	l.ticks++
}
