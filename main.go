package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/BlackAndWhite8/assets"
	_ "github.com/Zyko0/BlackAndWhite8/assets"
	"github.com/Zyko0/BlackAndWhite8/ui"

	"github.com/Zyko0/BlackAndWhite8/core"
	"github.com/Zyko0/BlackAndWhite8/graphics"
	"github.com/Zyko0/BlackAndWhite8/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	startMenu *ui.Menu
	pauseMenu *ui.Menu
	gameover  *ui.GameOver

	core     *core.Core
	renderer *graphics.Renderer
}

func New() *Game {
	start := ui.NewMenu(ui.GameTitle, ui.GameDescription, ui.GameStartKey)
	start.Initialize()
	start.Active = true

	pause := ui.NewMenu(ui.PauseTitle, ui.PauseDescription, ui.PauseResumeKey)
	pause.Initialize()

	gameover := ui.NewGameOver()
	gameover.Initialize()

	return &Game{
		startMenu: start,
		pauseMenu: pause,
		gameover:  gameover,

		renderer: graphics.NewRenderer(),
	}
}

func (g *Game) Update() error {
	// TODO: remove
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	// Start Menu
	if g.startMenu.Active {
		g.startMenu.Update()
		if !g.startMenu.Active {
			g.core = core.New()
			assets.StopMenuMusic()
			assets.SetMusic(assets.GameMusic)
			assets.ReplayMusic()
		}
		return nil
	}

	// Reset game
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.core = core.New()
		g.renderer.Loop = nil
		assets.StopMenuMusic()
		assets.SetMusic(assets.GameMusic)
		assets.ReplayMusic()
	}

	if g.core.IsOver() {
		if !g.core.IsPaused() {
			g.core.TogglePause()
		}
		g.gameover.Activate(g.core.GetStatistics())
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.core.TogglePause()
	}
	if g.core.IsPaused() {
		g.pauseMenu.Active = true
		g.pauseMenu.Update()
		return nil
	} else {
		g.pauseMenu.Active = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && g.renderer.Loop == nil {
		g.core.Loop()
		g.renderer.StartNewLoop(g.core.Player, g.core.Board.TileAt(g.core.Player.X, g.core.Player.Y))
		assets.SetMusic(assets.LoopMusic)
		assets.ReplayMusic()
	}

	if g.renderer.Loop == nil {
		assets.SetMusic(assets.GameMusic)
		g.core.Update()
	}
	g.renderer.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Start menu
	if g.startMenu.Active {
		g.startMenu.Draw(screen)
		return
	}

	g.renderer.RenderTiles(g.core.Board.Tiles)
	g.renderer.ClearEntities()
	if g.renderer.Loop != nil {
		g.renderer.RenderLoop(screen)
	} else {
		g.renderer.RenderEntities(g.core.Aoes, g.core.Projectiles)
		g.renderer.RenderPlayer(g.core.Player)
		g.renderer.Render(screen)
	}

	// Game over
	if g.core.IsOver() {
		g.gameover.Draw(screen)
		return
	}

	g.renderer.RenderHUD(screen, g.core.Player.HP, g.core.GetTime(), g.core.GetCompletion())

	// Pause menu
	if g.pauseMenu.Active {
		g.pauseMenu.Draw(screen)
		return
	}
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
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)

	assets.ResumeMenuMusic()

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
