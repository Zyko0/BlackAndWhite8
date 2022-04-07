package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	kindAoe        = 0
	kindProjectile = 1
)

func (r *Renderer) RenderAoes(aoes []*entity.Aoe) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	for i, aoe := range aoes {
		r := aoe.GetRect()
		size := r.Size()
		vertices, indices = AppendQuadVerticesIndices(
			vertices, indices,
			float32(r.Min.X), float32(r.Min.Y), float32(size.X), float32(size.Y),
			kindAoe, 0, 0, 0, i,
		)
	}

	r.offscreenEntities.DrawTrianglesShader(vertices, indices, assets.EntityShader, nil)
}

func (r *Renderer) RenderProjectiles(projectiles []*entity.Projectile) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	for i, proj := range projectiles {
		r := proj.GetRect()
		size := r.Size()
		vertices, indices = AppendQuadVerticesIndices(
			vertices, indices,
			float32(r.Min.X), float32(r.Min.Y), float32(size.X), float32(size.Y),
			kindProjectile, 0, 0, 0, i,
		)
	}

	r.offscreenEntities.DrawTrianglesShader(vertices, indices, assets.EntityShader, nil)
}
