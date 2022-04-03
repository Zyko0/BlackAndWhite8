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

func newBoard(rng *rand.Rand, difficulty Difficulty) *Board {
	size := 16
	if difficulty == DifficultyHard {
		size = 32
	}
	tileSize := logic.ScreenHeight / float32(size)

	tiles := make([][]*tile.Tile, size)
	kind := 0
	for y := range tiles {
		tiles[y] = make([]*tile.Tile, size)
		for x := range tiles[y] {
			tiles[y][x] = &tile.Tile{
				X:         float32(x) * tileSize,
				Y:         float32(y) * tileSize,
				W:         tileSize,
				H:         tileSize,
				KindIndex: kind, // tile.GetKind(rng.Intn(tile.MaxKind()), // TODO: re-enable this once we have enough variations
			}
			kind++
			if kind > tile.MaxKind() {
				kind = 0
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
