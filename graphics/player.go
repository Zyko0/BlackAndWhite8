package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderEntities(p *core.Player) {
	r.offscreenEntities.Clear()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	r.offscreenEntities.DrawImage(assets.PlayerIdleImage, op)
}
