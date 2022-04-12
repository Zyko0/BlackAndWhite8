package core

import (
	"math/rand"

	"github.com/Zyko0/BlackAndWhite8/assets/shape"
	"github.com/Zyko0/BlackAndWhite8/core/tile"
	"github.com/Zyko0/BlackAndWhite8/logic"
)

type Board struct {
	rng              *rand.Rand
	completed        int
	uncompletedTiles []*tile.Tile
	highlightedTile  *tile.Tile

	Size  int
	Tiles [][]*tile.Tile
}

func newBoard(rng *rand.Rand, shape *shape.Shape) *Board {
	const precompleted = 0.6

	tileSize := logic.ScreenHeight / float32(shape.Size())

	tiles := make([][]*tile.Tile, shape.Size())
	for y := range tiles {
		tiles[y] = make([]*tile.Tile, shape.Size())
		for x := range tiles[y] {
			kind := shape.At(x, y)
			if rng.Float32() > precompleted {
				kind = rng.Intn(tile.MaxKind())
			}
			tiles[y][x] = &tile.Tile{
				X:         x,
				Y:         y,
				W:         tileSize,
				H:         tileSize,
				KindIndex: kind,
				Completed: kind == shape.At(x, y),
			}
		}
	}

	return &Board{
		rng:              rng,
		uncompletedTiles: []*tile.Tile{},
		highlightedTile:  nil,

		Size:  shape.Size(),
		Tiles: tiles,
	}
}

func (b *Board) Update(s *shape.Shape) {
	b.completed = 0
	b.uncompletedTiles = b.uncompletedTiles[:0]

	// Update tiles
	for y := range b.Tiles {
		for _, t := range b.Tiles[y] {
			t.Update()
			if t.Completed {
				b.completed++
			} else {
				b.uncompletedTiles = append(b.uncompletedTiles, t)
			}
		}
	}
	// Unhighlight a completed tile
	if t := b.highlightedTile; t != nil && t.Completed {
		t.Highlighted = false
		b.highlightedTile = nil
	}
	// Refresh highlighted tile if necessary
	if t := b.highlightedTile; t == nil && len(b.uncompletedTiles) > 0 {
		index := b.rng.Intn(len(b.uncompletedTiles))
		b.highlightedTile = b.uncompletedTiles[index]
		b.highlightedTile.Highlighted = true
		b.uncompletedTiles[index] = b.uncompletedTiles[len(b.uncompletedTiles)-1]
		b.uncompletedTiles = b.uncompletedTiles[:len(b.uncompletedTiles)-1]
	}
}

func (b *Board) TileAt(x, y float32) *tile.Tile {
	vx := int(x / logic.ScreenHeight * float32(b.Size))
	vy := int(y / logic.ScreenHeight * float32(b.Size))
	if vx < 0 {
		vx = 0
	}
	if vy < 0 {
		vy = 0
	}
	if vx >= b.Size {
		vx = b.Size - 1
	}
	if vy >= b.Size {
		vy = b.Size - 1
	}

	return b.Tiles[vy][vx]
}
