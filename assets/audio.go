package assets

import (
	_ "embed"

	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume       = 0.5
	defaultGameMusicVolume = 0.7
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
	//go:embed sfx/dash.wav
	dashSoundBytes  []byte
	dashSoundPlayer *audio.Player

	//go:embed music/game.mp3
	gameMusicBytes  []byte
	gameMusicPlayer *audio.Player
)

func init() {
	var err error

	wavReader, err := wav.Decode(ctx, bytes.NewReader(hitSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	hitSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}
	hitSoundPlayer.SetVolume(defaultSFXVolume)

	wavReader, err = wav.Decode(ctx, bytes.NewReader(flipSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	flipSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}
	flipSoundPlayer.SetVolume(defaultSFXVolume)

	wavReader, err = wav.Decode(ctx, bytes.NewReader(flipFailSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	flipFailSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}

	wavReader, err = wav.Decode(ctx, bytes.NewReader(dashSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	dashSoundPlayer, err = ctx.NewPlayer(wavReader)
	if err != nil {
		log.Fatal(err)
	}

	mp3Reader, err := mp3.Decode(ctx, bytes.NewReader(gameMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader := audio.NewInfiniteLoop(mp3Reader, mp3Reader.Length())
	gameMusicPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	gameMusicPlayer.SetVolume(defaultGameMusicVolume)
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

func PlayDashSound() {
	dashSoundPlayer.Rewind()
	dashSoundPlayer.Play()
}

func ReplayGameMusic() {
	gameMusicPlayer.Rewind()
	if !gameMusicPlayer.IsPlaying() {
		gameMusicPlayer.Play()
	}
}

func ResumeInGameMusic() {
	if !gameMusicPlayer.IsPlaying() {
		gameMusicPlayer.Play()
	}
}

func StopInGameMusic() {
	if gameMusicPlayer.IsPlaying() {
		gameMusicPlayer.Pause()
	}
}
