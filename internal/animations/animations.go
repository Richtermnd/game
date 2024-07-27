package animations

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationState string

type SpriteHolder interface {
	GetSprites(state AnimationState) (sprite *ebiten.Image, numFramse int)
}

type AnimationSet struct {
	state          AnimationState
	holder         SpriteHolder
	currentSprite  *ebiten.Image
	animationSpeed int
	frameWidth     int
	frameHeight    int
	count          int
	numFrames      int
}

func New(
	spriteHolder SpriteHolder,
	state AnimationState,
	FPS float64,
	frameWidth, frameHeight int,
) *AnimationSet {
	// TODO: Allow slower animation speed (1 frame per 3 etc)
	if FPS > 1 || FPS < 0 {
		log.Fatal("FPS must be beetwen 0 and 1")
	}
	as := &AnimationSet{
		holder:         spriteHolder,
		animationSpeed: int(1 / FPS),
		frameWidth:     frameWidth,
		frameHeight:    frameHeight,
	}
	as.SetState(state)
	return as
}

func (as *AnimationSet) SetState(state AnimationState) {
	as.state = state
	as.currentSprite, as.numFrames = as.holder.GetSprites(state)
	as.count = 0
}

func (as *AnimationSet) NextFrame() *ebiten.Image {
	s := as.currentSprite              // short alias
	cf := as.count / as.animationSpeed // short alias
	x0, y0 := cf*as.frameWidth, 0
	x1, y1 := x0+as.frameWidth, as.frameHeight
	as.count = (as.count + 1) % (as.numFrames * as.animationSpeed)

	return s.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}
