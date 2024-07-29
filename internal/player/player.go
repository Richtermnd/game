package player

import (
	"github.com/Richtermnd/game/internal/animations"
	"github.com/Richtermnd/game/internal/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Pos   vector.Vector
	Dir   vector.Vector
	Speed float64
	anims *animations.AnimationSet
}

func NewPlayer(name string, x, y int) *Player {
	anims := animations.New(
		getSpriteHolder(name),
		IDLE,
		0.1, 24, 24,
	)
	return &Player{
		Pos:   vector.New(x, y),
		anims: anims,
		Speed: 5,
	}
}

func (p *Player) Update() error {
	p.handleInput()
	p.Pos = p.Pos.Add(p.Dir.Normalize().Scale(p.Speed))
	return nil
}

func (p *Player) handleInput() {
	p.handleMove()
}

func (p *Player) handleMove() {
	prevDir := p.Dir
	// Up and down
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Dir.Y = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Dir.Y = 1
	} else {
		p.Dir.Y = 0
	}

	// Left and right
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Dir.X = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Dir.X = 1
	} else {
		p.Dir.X = 0
	}

	l1, l2 := p.Dir.Len(), prevDir.Len()
	if l1 != l2 {
		if l1 == 0 {
			p.anims.SetState(IDLE)
		} else {
			p.anims.SetState(MOVE)
		}
	}
	if p.Dir.Len() != 0 && prevDir.Len() == 0 {
		p.anims.SetState(MOVE)
	} else if p.Dir.Len() == 0 && prevDir.Len() != 0 {
		p.anims.SetState(IDLE)
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
	ops.GeoM.Translate(float64(p.Pos.X), float64(p.Pos.Y))
	screen.DrawImage(frame, &ops)
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%+v %+v", p.Pos, p.Dir))
}
