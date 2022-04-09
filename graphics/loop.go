package graphics

import (
	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
)

type Loop struct {
	ticks uint64

	Done bool
}

func (l *Loop) Update() {

}

func (r *Renderer) StartNewLoop(p *core.Player, tile *tile.Tile) {
	r.Loop = &Loop{}
}
