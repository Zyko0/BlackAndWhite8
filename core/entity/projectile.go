package entity

import (
	"image"

	"github.com/Zyko0/BlackAndWhite8/logic"
)

const (
	projectileSpeed = 5.

	ProjectileRadius = float32(logic.ScreenHeight) / 64
)

type Projectile struct {
	X, Y float32

	dx, dy float32
	rect   image.Rectangle
}

func NewProjectile(x, y, dx, dy float32) *Projectile {
	return &Projectile{
		X:    x,
		Y:    y,
		dx:   dx,
		dy:   dy,
		rect: image.Rect(int(x), int(y), int(x+ProjectileRadius*2), int(y+ProjectileRadius*2)),
	}
}

func (p *Projectile) Update() {
	p.X += p.dx * projectileSpeed
	p.Y += p.dy * projectileSpeed
	p.rect = image.Rect(int(p.X), int(p.Y), int(p.X+ProjectileRadius*2), int(p.Y+ProjectileRadius*2))
}

// :)
func (p *Projectile) GetRect() image.Rectangle {
	return p.rect
}
