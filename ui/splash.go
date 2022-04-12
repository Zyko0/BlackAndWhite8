package ui

import (
	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	splashDisplayDuration = logic.TPS * 3
)

type Splash struct {
	ticks uint64

	geom   ebiten.GeoM
	colorm ebiten.ColorM

	Active bool
}

func NewSplash() *Splash {
	geom := ebiten.GeoM{}
	geom.Translate(0, 0)
	geom.Scale(
		float64(logic.ScreenWidth)/float64(assets.SplashScreenImage.Bounds().Max.X),
		float64(logic.ScreenHeight)/float64(assets.SplashScreenImage.Bounds().Max.Y),
	)
	return &Splash{
		ticks: 0,

		geom:   geom,
		colorm: ebiten.ColorM{},
	}
}

func (s *Splash) Update() {
	s.ticks++

	if s.ticks > splashDisplayDuration {
		s.Active = false
		return
	}
	if len(inpututil.AppendPressedKeys(nil)) > 0 || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.Active = false
	}

	d := float64(s.ticks) / float64(splashDisplayDuration)
	sc := (-(d * d) + d) * 4
	s.colorm.Reset()
	s.colorm.Scale(sc, sc, sc, 1.)
}

func (s *Splash) Draw(screen *ebiten.Image) {
	screen.DrawImage(assets.SplashScreenImage, &ebiten.DrawImageOptions{
		GeoM:   s.geom,
		Filter: ebiten.FilterLinear,
		ColorM: s.colorm,
	})
}
