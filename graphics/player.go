package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderPlayer(p *core.Player) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	var img *ebiten.Image
	switch {
	case p.InvulnDuration > 0:
		if p.InvulnDuration/10%2 == 0 {
			img = assets.PlayerInvuln0Image
		} else {
			img = assets.PlayerInvuln1Image
		}
	case p.DashDuration > 0:
		img = assets.PlayerDashImage
	default:
		img = assets.PlayerIdleImage
	}

	r.offscreenEntities.DrawImage(img, op)
}
