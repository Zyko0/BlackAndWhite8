package core

import (
	"math/rand"

	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

type Board struct {
	Size  int
	Tiles [][]*tile.Tile
}

func newBoard(rng *rand.Rand, size int) *Board {
	tileSize := logic.ScreenHeight / float32(size)

	tiles := make([][]*tile.Tile, size)
	for y := range tiles {
		tiles[y] = make([]*tile.Tile, size)
		for x := range tiles[y] {
			tiles[y][x] = &tile.Tile{
				X:         float32(x) * tileSize,
				Y:         float32(y) * tileSize,
				W:         tileSize,
				H:         tileSize,
				KindIndex: rng.Intn(tile.MaxKind()),
			}
		}
	}

	return &Board{
		Size:  size,
		Tiles: tiles,
	}
}

func (b *Board) Update() {
	for y := range b.Tiles {
		for x := range b.Tiles[y] {
			b.Tiles[y][x].Update()
		}
	}
}
