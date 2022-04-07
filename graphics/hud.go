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
		DurationX = 50
		DurationY = 100
	)

	text.Draw(
		screen,
		duration.String(),
		assets.HUDFontFace,
		DurationX, DurationY, color.White,
	)

	text.Draw(
		screen,
		fmt.Sprintf("%0.2f%%", completion*100),
		assets.HUDFontFace,
		DurationX, DurationY+75, color.White,
	)
}
