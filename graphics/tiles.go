package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderTiles(tiles [][]*tile.Tile) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	for y, row := range tiles {
		for x, tile := range row {
			index := y*len(row) + x
			outline := float32(0)
			if tile.Highlighted {
				outline = 1.
			}
			vertices, indices = AppendQuadVerticesIndices(
				vertices, indices,
				float32(tile.X)*tile.W, float32(tile.Y)*tile.H,
				tile.W, tile.H,
				float32(tile.GetKind().Method), float32(tile.GetKind().Arg), outline, r.highlightOffset, // TODO: /s/4/tile.RenderArg
				index,
			)
		}
	}

	r.offscreenBoard.DrawTrianglesShader(vertices, indices, assets.TileShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{},
	})
}
