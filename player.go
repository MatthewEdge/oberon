package oberon

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	speed    float32
	width    float32
	height   float32
	position rl.Vector2
	src      rl.Rectangle // Source on canvas
	dest     rl.Rectangle // Destination on canvas
	sprite   rl.Texture2D
}

func (p *Player) Init(x, y, width, height float32, texFile string) {
	p.sprite = rl.LoadTexture(texFile)
	p.width = width
	p.height = height
	p.speed = 1 // 0 - 1

	// TODO can these be replaced with a Vector2 position?
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
		p.position.Y -= p.speed
		p.dest.Y -= p.speed
	}
	if dir == DirDown {
		p.position.Y += p.speed
		p.dest.Y += p.speed
	}
	if dir == DirLeft {
		p.position.X -= p.speed
		p.dest.X -= p.speed
	}
	if dir == DirRight {
		p.position.X += p.speed
		p.dest.X += p.speed
	}
}

// GetCameraBounds returns the X/Y boundaries for the Player
func (p *Player) GetCameraBounds() (cx, cy float32) {
	cx = p.dest.X - (p.dest.Width / 2)
	cy = p.dest.Y - (p.dest.Height / 2)
	return cx, cy
}

// Draw draws the player
func (p *Player) Draw() {
	posStr := fmt.Sprintf("%f.2, %f.2", p.position.X, p.position.Y)
	rl.DrawText(posStr, int32(p.position.X), int32(p.position.Y-20), 14, rl.White)
	rl.DrawTexturePro(p.sprite, p.src, p.dest, rl.NewVector2(p.dest.Width, p.dest.Height), 0, rl.White)
}

// Close unloads the player textures
func (p *Player) Close() {
	rl.UnloadTexture(p.sprite)
}
