package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// Level
	bg rl.Color = rl.NewColor(0, 0, 0, 255)
)

type Game struct {
	width     int
	height    int
	renderFPS bool
}

func (g *Game) Init() {
	rl.InitWindow(640, 480, "Oberon")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func (g *Game) Input() {}

func (g *Game) Update() {}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bg)

	if g.renderFPS {
		rl.DrawFPS(5, 5)
	}

	rl.EndDrawing()
}

func (g *Game) ShouldClose() bool {
	return !rl.WindowShouldClose() // TODO
}

func main() {
	g := Game{
		width:     640,
		height:    480,
		renderFPS: true,
	}
	g.Init()

	for g.ShouldClose() {
		g.Input()
		g.Update()
		g.Draw()
	}

	rl.CloseWindow()
}
