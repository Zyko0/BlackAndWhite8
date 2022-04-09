package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridOffsetX = (logic.ScreenWidth - logic.ScreenHeight) / 2
)

type Renderer struct {
	offscreen         *ebiten.Image
	offscreenBoard    *ebiten.Image
	offscreenEntities *ebiten.Image

	highlightOffset float32

	Loop *Loop
}

func NewRenderer() *Renderer {
	return &Renderer{
		offscreen:         ebiten.NewImage(logic.ScreenWidth, logic.ScreenHeight),
		offscreenBoard:    ebiten.NewImage(logic.ScreenHeight, logic.ScreenHeight),
		offscreenEntities: ebiten.NewImage(logic.ScreenHeight, logic.ScreenHeight),
	}
}

func (r *Renderer) Update() {
	if r.Loop != nil {
		if r.Loop.Done {
			r.Loop = nil
		} else {
			r.Loop.Update()
		}
	}

	r.highlightOffset += 0.01
	if r.highlightOffset > 1 {
		r.highlightOffset = 0
	}
}

func (r *Renderer) ClearEntities() {
	r.offscreenEntities.Clear()
}

func (r *Renderer) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(GridOffsetX, 0)

	screen.DrawImage(r.offscreenBoard, op)
	screen.DrawImage(r.offscreenEntities, op)
}
