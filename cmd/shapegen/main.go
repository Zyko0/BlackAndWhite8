package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"strings"
)

var (
	//go:embed images
	images embed.FS

	palette    color.Palette
	grayImages []*image.Gray

	x16    [][][]int
	x16min []int
	x16max []int
)

func handleFile(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() || !strings.HasSuffix(d.Name(), ".png") {
		return nil
	}

	b, err := fs.ReadFile(images, path)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	grayImg := image.NewGray(img.Bounds())
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	grayImages = append(grayImages, grayImg)

	return nil
}

func init() {
	err := fs.WalkDir(images, ".", handleFile)
	if err != nil {
		log.Fatal(err)
	}

	// Palette
	const variationCount = 8
	for i := 0; i < variationCount; i++ {
		v := uint8(float64(256) / float64(variationCount) * float64(i))
		palette = append(palette, color.RGBA{v, v, v, 255})
	}
}

func main() {
	for _, img := range grayImages {
		valueCount := map[int]uint{}
		data := [][]int{}
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			data = append(data, []int{})
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				clr := img.At(x, y)
				// In palette
				i := palette.Index(clr)
				if _, ok := valueCount[i]; !ok {
					valueCount[i] = 1
				} else {
					valueCount[i]++
				}
				data[y] = append(data[y], i)
			}
		}
		fmt.Println("Unique colors count:", len(valueCount))

		min, max := 999, 0
		for v := range valueCount {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		fmt.Println("Max color value:", max)
		fmt.Println("Min color value:", min)

		x16 = append(x16, data)
		x16min = append(x16min, min)
		x16max = append(x16max, max)
	}

	Generate()
}
