package main

import (
	"oberon"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// Level
	bg rl.Color = rl.NewColor(0, 0, 0, 255)
)

type Game struct {
	winWidth  int32
	winHeight int32
	targetFPS int32
	renderFPS bool

	keymap oberon.Keymap

	player oberon.Player
}

func (g *Game) Init() {
	rl.InitWindow(g.winWidth, g.winHeight, "Oberon")
	rl.SetExitKey(0)
	rl.SetTargetFPS(g.targetFPS)

	// TODO load keymaps
	g.keymap.Init()

	g.player = oberon.Player{}
	g.player.Init(50, 50, 48, 48, "assets/chara_hero.png")
}

// Input handles translating player input to game actions
func (g *Game) Input() {
	// TODO the level should handle out-of-bounds
	// TODO but who handles collisions?
	if rl.IsKeyDown((g.keymap.Up)) {
		g.player.Move(oberon.DirUp)
	}
	if rl.IsKeyDown((g.keymap.Down)) {
		g.player.Move(oberon.DirDown)
	}
	if rl.IsKeyDown((g.keymap.Left)) {
		g.player.Move(oberon.DirLeft)
	}
	if rl.IsKeyDown((g.keymap.Right)) {
		g.player.Move(oberon.DirRight)
	}
}

func (g *Game) Update() {}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bg)

	// TODO dir
	g.player.Draw()

	if g.renderFPS {
		rl.DrawFPS(5, 5)
	}

	rl.EndDrawing()
}

func (g *Game) ShouldClose() bool {
	return !rl.WindowShouldClose() // TODO
}

// Close closes out resources
func (g *Game) Close() {
	g.player.Close()
	rl.CloseWindow()
}

func main() {
	g := Game{
		winWidth:  640,
		winHeight: 480,
		targetFPS: 60,
		renderFPS: true,
	}
	g.Init()
	defer g.Close()

	for g.ShouldClose() {
		g.Input()
		g.Update()
		g.Draw()
	}
}
