package ui

import (
	"math"
	"math/rand"

	"github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playersCount = 64
	playerSpeed  = 5.0
)

func sign(v float64) float64 {
	if v < 0 {
		return -1
	}
	return 1
}

type player struct {
	img    *ebiten.Image
	x, y   float64
	dx, dy float64
	angle  float64
}

func (p *player) Update() {
	p.angle += 0.01
	if p.angle > 1 {
		p.angle = 0
	}

	p.x += p.dx * playerSpeed
	p.y += p.dy * playerSpeed

	if p.x > logic.ScreenWidth {
		p.x = logic.ScreenWidth
		p.dx = -sign(p.dx) * rand.Float64()
	}
	if p.x < 0 {
		p.x = 0
		p.dx = -sign(p.dx) * rand.Float64()
	}
	if p.y > logic.ScreenHeight {
		p.y = logic.ScreenHeight
		p.dy = -sign(p.dy) * rand.Float64()
	}
	if p.y < 0 {
		p.y = 0
		p.dy = -sign(p.dy) * rand.Float64()
	}
}

type Background struct {
	bgImage *ebiten.Image

	players []*player
}

func NewBackground() *Background {
	return &Background{
		bgImage: ebiten.NewImage(logic.ScreenWidth, logic.ScreenHeight),

		players: make([]*player, playersCount),
	}
}

func (b *Background) Initialize() {
	for i := range b.players {
		var img *ebiten.Image
		switch rand.Intn(5) {
		case 0:
			img = assets.PlayerIdleImage
		case 1:
			img = assets.PlayerDashImage
		case 2:
			img = assets.PlayerInvuln0Image
		case 3:
			img = assets.PlayerInvuln1Image
		case 4:
			img = assets.PlayerLoopImage
		}
		p := &player{
			img:   img,
			x:     rand.Float64() * logic.ScreenWidth,
			y:     rand.Float64() * logic.ScreenHeight,
			dx:    rand.Float64()*2 - 1,
			dy:    rand.Float64()*2 - 1,
			angle: rand.Float64(),
		}
		b.players[i] = p
	}
}

func (b *Background) Update() {
	for _, p := range b.players {
		p.Update()
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.bgImage, nil)

	var geom ebiten.GeoM
	for _, p := range b.players {
		geom.Reset()
		geom.Translate(-float64(p.img.Bounds().Dx())/2, -float64(p.img.Bounds().Dy())/2)
		geom.Rotate(p.angle * math.Pi)
		geom.Translate(p.x, p.y)
		screen.DrawImage(p.img, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
	}
}
