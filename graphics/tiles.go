package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridOffsetX  = (logic.ScreenWidth - logic.ScreenHeight) / 2
	TileInterval = 16
)

func (r *Renderer) RenderTiles(screen *ebiten.Image, tiles [][]*tile.Tile) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	for y, row := range tiles {
		for x, tile := range row {
			index := y*len(row) + x
			vertices, indices = AppendQuadVerticesIndices(
				vertices, indices,
				GridOffsetX+tile.X, tile.Y,
				tile.W, tile.H,
				float32(tile.GetKind().Method), float32(tile.GetKind().Arg), 0, 0, // TODO: /s/4/tile.RenderArg
				index,
			)
		}
	}

	screen.DrawTrianglesShader(vertices, indices, assets.TileShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{},
	})
}
