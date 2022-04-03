package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileInterval = 16
)

func (r *Renderer) RenderTiles(tiles [][]*tile.Tile) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	for y, row := range tiles {
		for x, tile := range row {
			index := y*len(row) + x
			vertices, indices = AppendQuadVerticesIndices(
				vertices, indices,
				tile.X, tile.Y,
				tile.W, tile.H,
				float32(tile.GetKind().Method), float32(tile.GetKind().Arg), 0, 0, // TODO: /s/4/tile.RenderArg
				index,
			)
		}
	}

	r.offscreenBoard.DrawTrianglesShader(vertices, indices, assets.TileShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{},
	})
}
