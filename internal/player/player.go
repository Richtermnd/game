package player

import (
	"github.com/Richtermnd/game/internal/animations"
	"github.com/Richtermnd/game/internal/events"
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
	p := &Player{
		Pos:   vector.New(x, y),
		anims: anims,
		Speed: 5,
	}
	p.subscribeEvents()
	return p
}

func (p *Player) Update() error {
	p.Pos = p.Pos.Add(p.Dir.Normalize().Scale(p.Speed))
	return nil
}

func (p *Player) handleMove(e events.Event) {
	event := e.(*events.KeyboardEvent)

	if event.Key == ebiten.KeyW {
		switch event.EventType {
		case events.KeyPressed:
			p.Dir.Y--
		case events.KeyReleased:
			p.Dir.Y++
		}
	}

	if event.Key == ebiten.KeyS {
		switch event.EventType {
		case events.KeyPressed:
			p.Dir.Y++
		case events.KeyReleased:
			p.Dir.Y--
		}
	}

	if event.Key == ebiten.KeyA {
		switch event.EventType {
		case events.KeyPressed:
			p.Dir.X--
		case events.KeyReleased:
			p.Dir.X++
		}
	}

	if event.Key == ebiten.KeyD {
		switch event.EventType {
		case events.KeyPressed:
			p.Dir.X++
		case events.KeyReleased:
			p.Dir.X--
		}
	}
	if p.Dir.Len() == 0 {
		p.anims.SetState(IDLE)
	} else {
		p.anims.SetState(MOVE)
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

func (p *Player) subscribeEvents() {
	// Moving
	events.SubscribeKeyboard(p.handleMove, events.KeyPressed,
		ebiten.KeyW, ebiten.KeyA, ebiten.KeyS, ebiten.KeyD)
	events.SubscribeKeyboard(p.handleMove, events.KeyReleased,
		ebiten.KeyW, ebiten.KeyA, ebiten.KeyS, ebiten.KeyD)
}
