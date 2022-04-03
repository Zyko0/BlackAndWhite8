package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridOffsetX = (logic.ScreenWidth - logic.ScreenHeight) / 2
)

type Renderer struct {
	offscreenBoard    *ebiten.Image
	offscreenEntities *ebiten.Image
}

func NewRenderer() *Renderer {
	return &Renderer{
		offscreenBoard:    ebiten.NewImage(logic.ScreenHeight, logic.ScreenHeight),
		offscreenEntities: ebiten.NewImage(logic.ScreenHeight, logic.ScreenHeight),
	}
}

func (r *Renderer) Update() {

}

func (r *Renderer) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(GridOffsetX, 0)

	screen.DrawImage(r.offscreenBoard, op)
	screen.DrawImage(r.offscreenEntities, op)
}
