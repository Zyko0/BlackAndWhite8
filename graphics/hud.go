package graphics

import (
	"fmt"
	"image/color"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (r *Renderer) RenderHUD(screen *ebiten.Image, playerHP int, duration time.Duration, completion, difficulty float64) {
	const (
		DurationX   = 50
		DurationY   = 100
		CompletionX = 50
		CompletionY = 125
		HpX         = 50
		HpY         = 150
		DifficultyX = 50
		DifficultyY = 160
	)

	ms := duration.Milliseconds()
	sec := ms / 1000 % 60
	ms = ms % 1000 / 10
	min := int(duration.Minutes())
	// Duration
	text.Draw(
		screen,
		fmt.Sprintf("%02d:%02d:%02d", min, sec, ms),
		assets.DefaultFontFace,
		DurationX, DurationY, color.White,
	)
	// Completion
	text.Draw(
		screen,
		fmt.Sprintf("%0.2f%%", completion*100),
		assets.DefaultFontFace,
		CompletionX, CompletionY, color.White,
	)
	// Player HP
	str := ""
	for i := 0; i < playerHP; i++ {
		str += "♥"
	}
	text.Draw(
		screen,
		str,
		assets.DefaultFontFace,
		HpX, HpY, color.White,
	)
	// Difficulty
	vertices, indices := AppendQuadVerticesIndices(
		nil, nil,
		DifficultyX, DifficultyY,
		float32(110*difficulty), 16,
		1, 1, 1, 1, 0,
	)
	screen.DrawTriangles(vertices, indices, BrushImage, nil)
}
