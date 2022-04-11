package assets

import (
	_ "embed"

	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume       = 0.5
	defaultGameMusicVolume = 0.4
	defaultMainMenuVolume  = 1.0
)

var (
	ctx = audio.NewContext(44100)

	//go:embed sfx/hit.wav
	hitSoundBytes  []byte
	hitSoundPlayer *audio.Player
	//go:embed sfx/flip.wav
	flipSoundBytes  []byte
	flipSoundPlayer *audio.Player
	//go:embed sfx/flipfail.wav
	flipFailSoundBytes  []byte
	flipFailSoundPlayer *audio.Player
)

func init() {
	var err error

	reader, err := wav.Decode(ctx, bytes.NewReader(hitSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	hitSoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	hitSoundPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(flipSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	flipSoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	flipSoundPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(flipFailSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	flipFailSoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
}

func PlayHitSound() {
	hitSoundPlayer.Rewind()
	hitSoundPlayer.Play()
}

func PlayFlipSound() {
	flipSoundPlayer.Rewind()
	flipSoundPlayer.Play()
}

func PlayFlipFailSound() {
	flipFailSoundPlayer.Rewind()
	flipFailSoundPlayer.Play()
}
