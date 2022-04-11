package ui

import (
	"image/color"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/graphics"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	GameTitle       = "Oprelo"
	GameDescription = "Press <Enter> to start"
	GameStartKey    = ebiten.KeyEnter

	PauseTitle       = "Pause"
	PauseDescription = "Press <P> to resume"
	PauseResumeKey   = ebiten.KeyP
)

type Menu struct {
	title, description string
	key                ebiten.Key
	bgImage            *ebiten.Image

	Active bool
}

func NewMenu(title, desc string, stopKey ebiten.Key) *Menu {
	return &Menu{
		title:       title,
		description: desc,
		key:         stopKey,
		bgImage:     ebiten.NewImage(logic.ScreenHeight/2, logic.ScreenHeight/2),

		Active: false,
	}
}

func (m *Menu) Initialize() {
	const (
		borderWidth       = 4
		cardWidth         = logic.ScreenHeight / 2
		noBorderCardWidth = logic.ScreenHeight/2 - borderWidth*2
	)

	m.bgImage.Fill(color.White)

	vertices, indices := graphics.AppendQuadVerticesIndices(
		nil, nil,
		borderWidth, borderWidth,
		noBorderCardWidth, noBorderCardWidth,
		0, 0, 0, 1, 0,
	)
	m.bgImage.DrawTriangles(vertices, indices, graphics.BrushImage, nil)
	// Title
	rect := text.BoundString(assets.DefaultFontFace, m.title)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(m.bgImage.Bounds().Dx()/2-rect.Max.X/2),
		float64(36),
	)
	text.DrawWithOptions(m.bgImage, m.title, assets.DefaultFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Exit text
	rect = text.BoundString(assets.DefaultSmallFontFace, m.description)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(m.bgImage.Bounds().Dx()/2-rect.Max.X/2),
		float64(36*2),
	)
	text.DrawWithOptions(m.bgImage, m.description, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(m.key) {
		m.Active = false
		return
	}

	x, y := ebiten.CursorPosition()
	// Volume management
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if x > logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2+96 &&
			x < logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2+96+128 &&
			y > logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2+36*3 &&
			y < logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2+36*3+16 {
			assets.SetMusicVolume(float64(x-logic.ScreenWidth/2+m.bgImage.Bounds().Dx()/2-96) / 128)
		} else if x > logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2+96 &&
			x < logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2+96+128 &&
			y > logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2+36*4 &&
			y < logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2+36*4+16 {
			assets.SetSFXVolume(float64(x-logic.ScreenWidth/2+m.bgImage.Bounds().Dx()/2-96) / 128)
		}
	}
}

func (m *Menu) drawVolume(screen *ebiten.Image, y, volume float64, title string) {
	rect := text.BoundString(assets.DefaultSmallFontFace, title)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2)+24,
		float64(logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2+rect.Dy())+y,
	)
	text.DrawWithOptions(screen, title, assets.DefaultSmallFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Volume
	vertices, indices := graphics.AppendQuadVerticesIndices(
		nil, nil,
		float32(logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2)+96,
		float32(logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2)+float32(y),
		float32(volume)*128, 16,
		1, 1, 1, 1, 0,
	)
	screen.DrawTriangles(vertices, indices, graphics.BrushImage, nil)
}

func (m *Menu) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(logic.ScreenWidth/2-m.bgImage.Bounds().Dx()/2),
		float64(logic.ScreenHeight/2-m.bgImage.Bounds().Dy()/2),
	)
	screen.DrawImage(m.bgImage, op)

	m.drawVolume(screen, 36*3, assets.GetMusicVolume(), "Music")
	m.drawVolume(screen, 36*4, assets.GetSFXVolume(), "Effects")
}
