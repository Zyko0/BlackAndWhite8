package entity

import (
	"image"
	"math"

	"github.com/Zyko0/BlackAndWhite8/logic"
)

const (
	aoeDuration     = logic.TPS * 1
	AoeEndTick      = logic.TPS * 1.25
	AoeDefaultWidth = float32(logic.ScreenHeight) / 16
)

type Aoe struct {
	ticks int

	x, y   float32
	width  float32
	length float32
}

func NewAoe(x, y, w float32) *Aoe {
	return &Aoe{
		ticks: 0,

		x:      x,
		y:      y,
		width:  w,
		length: 0,
	}
}

func (a *Aoe) Update() {
	if a.ticks <= aoeDuration {
		t := float64(a.ticks) / aoeDuration
		if t > 0 {
			t = math.Pow(2, 6*(t-1))
		}
		a.length = float32(t) * logic.ScreenHeight
	}

	a.ticks++
}

func (a *Aoe) IsOver() bool {
	return a.ticks > AoeEndTick
}

// :)
func (a *Aoe) GetRect() image.Rectangle {
	var x0, y0, w, h float32

	switch {
	case a.x == 0:
		x0, y0, w, h = a.x, a.y, a.length, a.width
	case a.x == logic.ScreenHeight:
		x0, y0, w, h = a.x-a.length, a.y, a.length, a.width
	case a.y == 0:
		x0, y0, w, h = a.x, a.y, a.width, a.length
	case a.y == logic.ScreenHeight:
		x0, y0, w, h = a.x, a.y-a.length, a.width, a.length
	}

	return image.Rectangle{
		Min: image.Point{
			X: int(x0),
			Y: int(y0),
		},
		Max: image.Point{
			X: int(x0 + w),
			Y: int(y0 + h),
		},
	}
}
