package utils

import (
	"image"
	"math"
)

func GetKnockbackVector(player, entity image.Rectangle) (float32, float32) {
	centerEntity := image.Point{
		X: (entity.Max.X-entity.Min.X)/2 + entity.Min.X,
		Y: (entity.Max.Y-entity.Min.Y)/2 + entity.Min.Y,
	}
	centerPlayer := image.Point{
		X: (player.Max.X-player.Min.X)/2 + player.Min.X,
		Y: (player.Max.Y-player.Min.Y)/2 + player.Min.Y,
	}
	dx := float32(centerPlayer.X - centerEntity.X)
	dy := float32(centerPlayer.Y - centerEntity.Y)

	// normalize
	length := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	dx /= length
	dy /= length

	return dx, dy
}
