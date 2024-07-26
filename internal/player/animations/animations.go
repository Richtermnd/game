package animations

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	frameSize      = 24
	animationSpeed = 8
)

type AnimationState string

// AnimationState
const (
	AVOID      AnimationState = "avoid"
	BITE       AnimationState = "bite"
	DASH       AnimationState = "dash"
	DEAD       AnimationState = "dead"
	HURT       AnimationState = "hurt"
	IDLE       AnimationState = "idle"
	JUMP       AnimationState = "jump"
	KICK       AnimationState = "kick"
	MOVE       AnimationState = "move"
	SCAN       AnimationState = "scan"
	GHOST_IDLE AnimationState = "ghost_idle"
	GHOST_MOVE AnimationState = "ghost_move"
)

type AnimationSet struct {
	state         AnimationState
	holder        *SpriteHolder
	currentSprite *ebiten.Image
	currentFrame  int
	numFrames     int
}

func New(name string) *AnimationSet {
	return &AnimationSet{
		state:  IDLE,
		holder: getSpriteHolder(name),
	}
}

func (as *AnimationSet) SetState(state AnimationState) {
	as.state = state
	as.currentFrame = 0
	switch state {
	case IDLE:
		as.numFrames = 3
		as.currentSprite = as.holder.Idle
	case MOVE:
		as.numFrames = 6
		as.currentSprite = as.holder.Move
	default:
		log.Fatalf("Unimplemeted state: %v", state)
	}
}

func (as *AnimationSet) NextFrame() *ebiten.Image {
	s := as.currentSprite                  // short alias
	cf := as.currentFrame / animationSpeed // short alias
	x0, y0 := cf*frameSize, 0
	x1, y1 := x0+frameSize, frameSize
	as.currentFrame = (as.currentFrame + 1) % (as.numFrames * animationSpeed)

	return s.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

// // incrementCurrentFrame
// // Move it to another method to make it not so cursed
// // TODO: Make it not so cursed
// func (as *AnimationSet) incrementCurrentFrame() {
// 	as.currentFrame++
// }
