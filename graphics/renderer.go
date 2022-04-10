package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridOffsetX = (logic.ScreenWidth - logic.ScreenHeight) / 2
)

type Renderer struct {
	offscreenBoard    *ebiten.Image
	offscreenEntities *ebiten.Image

	highlightOffset float32

	Loop *Loop
}

func NewRenderer() *Renderer {
	return &Renderer{
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

func (r *Renderer) StartNewLoop(p *core.Player, tile *tile.Tile) {
	r.Loop = &Loop{
		tx:           float64(p.X + core.PlayerSize/2),
		ty:           float64(p.Y + core.PlayerSize/2),
		currentScale: 1,
		bgScale:      64,
	}
}

func (r *Renderer) RenderLoop(screen *ebiten.Image) {
	const tileSize = logic.ScreenHeight / 16

	// Zoom centered on player tile
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(r.Loop.currentScale, r.Loop.currentScale)
	op.GeoM.Translate(
		GridOffsetX+(r.Loop.tx)*(1-r.Loop.currentScale),
		(r.Loop.ty)*(1-r.Loop.currentScale),
	)
	screen.DrawImage(r.offscreenBoard, op)
	// Void quad
	op = &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeClear,
	}
	op.GeoM.Scale(tileSize*r.Loop.currentScale, tileSize*r.Loop.currentScale)
	op.GeoM.Translate(
		GridOffsetX+(r.Loop.tx)-(tileSize/2.)*r.Loop.currentScale,
		(r.Loop.ty)-(tileSize/2.)*r.Loop.currentScale,
	)
	screen.DrawImage(brushImage, op)
	// Background
	vertices, indices := AppendQuadVerticesIndices(
		nil, nil,
		0, 0,
		logic.ScreenWidth, logic.ScreenHeight,
		1, 1, 1, 1, 0,
	)
	for i := range vertices {
		vertices[i].SrcX *= logic.ScreenWidth
		vertices[i].SrcY *= logic.ScreenHeight
	}
	screen.DrawTrianglesShader(vertices, indices, assets.GridShader, &ebiten.DrawTrianglesShaderOptions{
		CompositeMode: ebiten.CompositeModeDestinationOver,
		Uniforms: map[string]interface{}{
			"Scale": float32(r.Loop.bgScale),
			"Origin": []float32{
				float32((r.Loop.tx + GridOffsetX) / logic.ScreenHeight * r.Loop.bgScale),
				float32(r.Loop.ty / logic.ScreenHeight * r.Loop.bgScale),
			},
		},
		Images: [4]*ebiten.Image{
			r.offscreenBoard,
		},
	})
	// Player
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		GridOffsetX+(r.Loop.tx)-float64(core.PlayerSize)/2,
		(r.Loop.ty)-float64(core.PlayerSize)/2,
	)
	screen.DrawImage(assets.PlayerLoopImage, op)
}
