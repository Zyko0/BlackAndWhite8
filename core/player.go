package core

import (
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	MoveSpeed     = float32(6.5)
	DashSpeed     = float32(20)
	PlayerSize    = float32(32)
	DashCooldown  = 20
	DashDuration  = 10
	InvulnTime    = logic.TPS
	KnockbackTime = logic.TPS * 0.2 // TODO: tune
)

type Player struct {
	X, Y float32
	HP   int

	intentX, intentY int
	intentDash       bool
	intentFlip       bool
	intentLoop       bool

	DashCD            int
	DashDuration      int
	InvulnDuration    int
	KnockbackDuration int
}

func newPlayer() *Player {
	return &Player{
		X:  logic.ScreenHeight / 2,
		Y:  logic.ScreenHeight / 2,
		HP: 5,
	}
}

func (p *Player) Update() {
	p.intentX = 0
	p.intentY = 0
	p.intentDash = false
	p.intentFlip = false
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
		p.intentFlip = true
	}
}
