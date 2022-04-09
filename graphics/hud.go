package graphics

import (
	"fmt"
	"image/color"
	"time"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (r *Renderer) RenderHUD(screen *ebiten.Image, playerHP int, duration time.Duration, completion float64) {
	const (
		DurationX   = 50
		DurationY   = 100
		CompletionX = 50
		CompletionY = 125
		HpX         = 50
		HpY         = 150
	)

	d := ""
	ms := duration.Milliseconds()
	sec := ms / 1000 % 60
	ms = ms % 1000 / 10
	min := int(duration.Minutes())
	if min > 0 {
		d = fmt.Sprintf("%02d:%02d:%02d", min, sec, ms)
	} else {
		d = fmt.Sprintf("%02d:%02d", sec, ms)
	}
	// Duration
	text.Draw(
		screen,
		d,
		assets.HUDFontFace,
		DurationX, DurationY, color.White,
	)
	// Completion
	text.Draw(
		screen,
		fmt.Sprintf("%0.2f%%", completion*100),
		assets.HUDFontFace,
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
		assets.HUDFontFace,
		HpX, HpY, color.White,
	)
}
