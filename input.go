package oberon

import rl "github.com/gen2brain/raylib-go/raylib"

// Keymap represents the mapping from a logical key to the actual key ID. Allows the player
// to remap these actions
type Keymap struct {
	Up    int32
	Down  int32
	Left  int32
	Right int32
	Pause int32
	Exit  int32
}

// TODO defaults
func (k *Keymap) Init() {
	k.Up = rl.KeyW
	k.Down = rl.KeyS
	k.Left = rl.KeyA
	k.Right = rl.KeyD
	k.Pause = rl.KeyP
	k.Exit = rl.KeyEscape
}
