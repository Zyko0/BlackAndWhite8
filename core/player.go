package core

import (
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	MoveSpeed  = float32(5.0)
	PlayerSize = float32(32.0)
)

type Player struct {
	X, Y float32

	intentX, intentY float32
	intentDash       bool
	intentFlipUp     bool
	intentFlipDown   bool
	intentLoop       bool
}

func newPlayer() *Player {
	return &Player{
		X: logic.ScreenHeight / 2,
		Y: logic.ScreenHeight / 2,
	}
}

func (p *Player) Update() {
	p.intentX = 0
	p.intentY = 0
	p.intentDash = false
	p.intentFlipUp = false
	p.intentFlipDown = false
	p.intentLoop = false

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.intentY = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.intentY = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.intentX = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.intentX = 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.intentDash = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		p.intentLoop = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		p.intentFlipUp = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		p.intentFlipDown = true
	}
}
