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
	musicVolume = float64(1)
	sfxVolume   = float64(1)

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
	//go:embed music/menu.mp3
	menuMusicBytes  []byte
	menuMusicPlayer *audio.Player
	//go:embed music/loop.mp3
	loopMusicBytes  []byte
	loopMusicPlayer *audio.Player
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

	mp3Reader, err = mp3.Decode(ctx, bytes.NewReader(menuMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoop(mp3Reader, mp3Reader.Length())
	menuMusicPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}

	mp3Reader, err = mp3.Decode(ctx, bytes.NewReader(loopMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoop(mp3Reader, mp3Reader.Length())
	loopMusicPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSFXVolume() float64 {
	return sfxVolume
}

func SetSFXVolume(v float64) {
	sfxVolume = v
	hitSoundPlayer.SetVolume(v * defaultSFXVolume)
	flipSoundPlayer.SetVolume(v * defaultSFXVolume)
	flipFailSoundPlayer.SetVolume(v * defaultSFXVolume)
	dashSoundPlayer.SetVolume(v * defaultSFXVolume)
}

func GetMusicVolume() float64 {
	return musicVolume
}

func SetMusicVolume(v float64) {
	musicVolume = v
	gameMusicPlayer.SetVolume(v * defaultGameMusicVolume)
	menuMusicPlayer.SetVolume(v * defaultMainMenuVolume)
	loopMusicPlayer.SetVolume(v * defaultMainMenuVolume)
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
