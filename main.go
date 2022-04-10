package main

import (
	"errors"
	"fmt"

	_ "github.com/Zyko0/BlackAndWhite8/assets"

	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/graphics"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	pause      bool
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
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.pause = !g.pause
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.core.Loop()
		g.renderer.StartNewLoop(g.core.Player, g.core.Board.TileAt(g.core.Player.X, g.core.Player.Y))
	}

	// Reset game
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.core = core.New(g.difficulty)
		g.renderer.Loop = nil
	}

	if !g.pause && g.renderer.Loop == nil {
		g.core.Update()
	}
	g.renderer.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.RenderTiles(g.core.Board.Tiles)
	g.renderer.ClearEntities()
	if g.renderer.Loop != nil {
		g.renderer.RenderLoop(screen)
	} else {
		g.renderer.RenderEntities(g.core.Aoes, g.core.Projectiles)
		g.renderer.RenderPlayer(g.core.Player)
		g.renderer.Render(screen)
	}
	g.renderer.RenderHUD(screen, g.core.Player.HP, g.core.GetTime(), g.core.GetCompletion())
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
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOn) // TODO: do
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
