package assets

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

var (
	HUDFontFace font.Face
)

func init() {
	pfont, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		log.Fatal(err)
	}
	HUDFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 36,
	})
}
