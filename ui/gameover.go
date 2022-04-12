package ui

import (
	"fmt"
	"image/color"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/graphics"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameOver struct {
	title      string
	completion string
	duration   string
	loopCount  string
	hp         string

	bgImage *ebiten.Image
}

func NewGameOver() *GameOver {
	return &GameOver{
		bgImage: ebiten.NewImage(logic.ScreenHeight/2, logic.ScreenHeight/2),
	}
}

func (g *GameOver) Initialize() {
	const (
		borderWidth       = 4
		cardWidth         = logic.ScreenHeight / 2
		noBorderCardWidth = logic.ScreenHeight/2 - borderWidth*2
	)

	g.bgImage.Fill(color.White)

	vertices, indices := graphics.AppendQuadVerticesIndices(
		nil, nil,
		borderWidth, borderWidth,
		noBorderCardWidth, noBorderCardWidth,
		0, 0, 0, 1, 0,
	)
	g.bgImage.DrawTriangles(vertices, indices, graphics.BrushImage, nil)
}

func (g *GameOver) Activate(stats core.Statistics) {
	win := stats.HP > 0
	if win {
		g.title = "Win !"
	} else {
		g.title = "Game over"
	}

	g.completion = fmt.Sprintf("Completion %.2f%%", stats.Completion*100)
	ms := stats.Duration.Milliseconds()
	sec := ms / 1000 % 60
	ms = ms % 1000 / 10
	min := int(stats.Duration.Minutes())
	g.duration = fmt.Sprintf("Survived %02d:%02d:%02d", min, sec, ms)
	g.loopCount = fmt.Sprintf("Loops: %d", stats.LoopCount)
	g.hp = "HP: "
	for i := 0; i < stats.HP; i++ {
		g.hp += "♥"
	}
}

func (g *GameOver) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(logic.ScreenWidth/2-g.bgImage.Bounds().Dx()/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2),
	)
	screen.DrawImage(g.bgImage, op)

	// Title
	rect := text.BoundString(assets.DefaultFontFace, g.title)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36,
	)
	text.DrawWithOptions(screen, g.title, assets.DefaultFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Completion text
	rect = text.BoundString(assets.DefaultSmallFontFace, g.completion)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*2,
	)
	text.DrawWithOptions(screen, g.completion, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Duration text
	rect = text.BoundString(assets.DefaultSmallFontFace, g.duration)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*3,
	)
	text.DrawWithOptions(screen, g.duration, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Loop count text
	rect = text.BoundString(assets.DefaultSmallFontFace, g.loopCount)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*4,
	)
	text.DrawWithOptions(screen, g.loopCount, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// HP text
	rect = text.BoundString(assets.DefaultSmallFontFace, g.hp)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*5,
	)
	text.DrawWithOptions(screen, g.hp, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Press R to restart
	const str = "Press <Backspace> to restart"

	rect = text.BoundString(assets.DefaultSmallFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-rect.Max.X/2),
		float64(logic.ScreenHeight/2-g.bgImage.Bounds().Dy()/2+rect.Dy())+36*6.5,
	)
	text.DrawWithOptions(screen, str, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}
