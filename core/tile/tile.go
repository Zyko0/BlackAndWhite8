package tile

type Tile struct {
	X, Y int
	W, H float32
	// kind
	KindIndex   int
	Completed   bool
	Highlighted bool
}

func (t *Tile) Update() {

}

func (t *Tile) GetKind() Kind {
	return kinds[t.KindIndex]
}

func (t *Tile) FlipDown() {
	t.KindIndex--
	if t.KindIndex < 0 {
		t.KindIndex = len(kinds) - 1
	}
}

func (t *Tile) FlipUp() {
	t.KindIndex++
	if t.KindIndex > len(kinds)-1 {
		t.KindIndex = 0
	}
}
