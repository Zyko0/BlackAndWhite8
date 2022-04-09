package assets

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/tile.kage
	tileShaderSrc []byte
	//go:embed shaders/entity.kage
	entityShaderSrc []byte

	TileShader      *ebiten.Shader
	EntityShader    *ebiten.Shader
)

func init() {
	var err error

	TileShader, err = ebiten.NewShader(tileShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	EntityShader, err = ebiten.NewShader(entityShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
