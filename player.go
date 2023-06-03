package oberon

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	speed  float32
	width  float32
	height float32
	src    rl.Rectangle // Source on canvas
	dest   rl.Rectangle // Destination on canvas
	sprite rl.Texture2D
}

func (p *Player) Init(x, y, width, height float32, texFile string) {
	p.width = width
	p.height = height
	p.speed = 3 // TODO
	p.sprite = rl.LoadTexture(texFile)

	p.src = rl.NewRectangle(0, 0, p.width, p.height)
	p.dest = rl.NewRectangle(x, y, p.width, p.height)
}

const (
	DirUp = iota
	DirDown
	DirLeft
	DirRight
)

// Move sets the player to move in the given direction
func (p *Player) Move(dir int32) {
	if dir == DirUp {
		p.dest.Y -= p.speed
	}
	if dir == DirDown {
		p.dest.Y += p.speed
	}
	if dir == DirLeft {
		p.dest.X -= p.speed
	}
	if dir == DirRight {
		p.dest.X += p.speed
	}
}

// Draw draws the player
func (p *Player) Draw() {
	rl.DrawTexturePro(p.sprite, p.src, p.dest, rl.NewVector2(p.dest.Width, p.dest.Height), 0, rl.White)
}

// Close unloads the player textures
func (p *Player) Close() {
	rl.UnloadTexture(p.sprite)
}
