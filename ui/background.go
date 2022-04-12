package ui

import (
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	bgImage *ebiten.Image
}

func NewBackground() *Background {
	return &Background{
		bgImage: ebiten.NewImage(logic.ScreenWidth, logic.ScreenHeight),
	}
}

func (b *Background) Initialize() {

}

func (b *Background) Update() {
	
}