package player

import (
	"github.com/Richtermnd/game/internal/animations"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y  int
	anims *animations.AnimationSet
}

func NewPlayer(name string, x, y int) *Player {
	anims := animations.New(
		getSpriteHolder(name),
		IDLE,
		0.1, 24, 24,
	)
	anims.SetState(IDLE)
	return &Player{
		X:     x,
		Y:     y,
		anims: anims,
	}
}

func (p *Player) Layout() int {
	return 10
}

func (p *Player) Draw(screen *ebiten.Image) {
	frame := p.anims.NextFrame()
	w, h := frame.Bounds().Dx(), frame.Bounds().Dy()
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Scale(3, 3)
	ops.GeoM.Translate(-float64(w/2), -float64(h/2))
	ops.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(frame, &ops)
}
