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
	zoom      float32
	paused    bool

	keymap oberon.Keymap

	player oberon.Player
	cam    rl.Camera2D
}

func (g *Game) Init() {
	// Must init Textures and Sounds AFTER Window/Audio init
	rl.InitWindow(g.winWidth, g.winHeight, "Oberon")
	rl.SetExitKey(0)
	rl.SetTargetFPS(g.targetFPS)

	g.zoom = 1.0

	// TODO load keymaps
	g.keymap.Init()

	g.player = oberon.Player{}
	g.player.Init(50, 50, 48, 48, "assets/chara_hero.png")

	g.cam = rl.NewCamera2D(
		rl.NewVector2(float32(g.winWidth/2), float32(g.winHeight/2)),
		rl.NewVector2(g.player.GetCameraBounds()),
		0.0,
		g.zoom)
}

func (g *Game) Update() {
	// TODO the level should handle out-of-bounds
	// TODO but who handles collisions?
	if rl.IsKeyDown(g.keymap.Up) {
		g.player.Move(oberon.DirUp)
	}
	if rl.IsKeyDown(g.keymap.Down) {
		g.player.Move(oberon.DirDown)
	}
	if rl.IsKeyDown(g.keymap.Left) {
		g.player.Move(oberon.DirLeft)
	}
	if rl.IsKeyDown(g.keymap.Right) {
		g.player.Move(oberon.DirRight)
	}

	// Update camera location
	g.cam.Target = rl.NewVector2(g.player.GetCameraBounds())
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(bg)
	rl.BeginMode2D(g.cam)

	// TODO swap to Map.Draw()
	g.player.Draw()

	// TODO move to top of player view
	if g.renderFPS {
		rl.DrawFPS(5, 5)
	}

	rl.EndMode2D()
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
		g.Update()
		g.Draw()
	}
}
