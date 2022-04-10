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
	//go:embed shaders/grid.kage
	gridShaderSrc []byte

	TileShader   *ebiten.Shader
	EntityShader *ebiten.Shader
	GridShader   *ebiten.Shader
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

	GridShader, err = ebiten.NewShader(gridShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
