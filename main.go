package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/graphics"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	difficulty core.Difficulty

	core     *core.Core
	renderer *graphics.Renderer
}

func New() *Game {
	return &Game{
		difficulty: core.DifficultyNormal,
		core:       core.New(core.DifficultyNormal),
		renderer:   graphics.NewRenderer(),
	}
}

func (g *Game) Update() error {
	// TODO: remove
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	// Reset game
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.core = core.New(core.DifficultyHard) // g.difficulty)
	}

	g.core.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.RenderTiles(g.core.Board.Tiles)
	g.renderer.RenderEntities(g.core.Player)
	g.renderer.Render(screen)
	// Debug
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f - FPS: %0.2f",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
