package assets

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/tile.kage
	tileShaderSrc []byte

	TileShader *ebiten.Shader
)

func init() {
	var err error

	TileShader, err = ebiten.NewShader(tileShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
