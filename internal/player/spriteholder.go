package player

import (
	"log"
	"os"
	"sync"

	"github.com/Richtermnd/game/internal/animations"
	"github.com/hajimehoshi/ebiten/v2"
)

// alias
type animState = animations.AnimationState

// States
const (
	AVOID      animState = "avoid"
	BITE       animState = "bite"
	DASH       animState = "dash"
	DEAD       animState = "dead"
	HURT       animState = "hurt"
	IDLE       animState = "idle"
	JUMP       animState = "jump"
	KICK       animState = "kick"
	MOVE       animState = "move"
	SCAN       animState = "scan"
	GHOST_IDLE animState = "ghost_idle"
	GHOST_MOVE animState = "ghost_move"
)

var (
	holdersMap   map[string]*spriteHolder = map[string]*spriteHolder{}
	mu           sync.Mutex
	SpriteHolder = new(spriteHolder)
)

type spriteHolder struct {
	Avoid     *ebiten.Image
	Bite      *ebiten.Image
	Dash      *ebiten.Image
	Dead      *ebiten.Image
	Hurt      *ebiten.Image
	Idle      *ebiten.Image
	Jump      *ebiten.Image
	Kick      *ebiten.Image
	Move      *ebiten.Image
	Scan      *ebiten.Image
	GhostIdle *ebiten.Image
	GhostMove *ebiten.Image
}

func (s *spriteHolder) GetSprites(state animState) (*ebiten.Image, int) {
	switch state {
	case IDLE:
		return s.Idle, 3
	case MOVE:
		return s.Move, 6
	default:
		log.Fatalf("Unimplemented state: %s", state)
		return nil, 0
	}
}

const spriteFolder = "sprites/players"

func init() {
	dinos, err := os.ReadDir(spriteFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, dino := range dinos {
		uploadDino(dino.Name())
	}
}

func getSpriteHolder(name string) *spriteHolder {
	mu.Lock()
	defer mu.Unlock()
	holder, ok := holdersMap[name]
	if !ok {
		log.Fatal("Unknown sprite name")
	}
	return holder
}

func uploadDino(dino string) {
	path := spriteFolder + "/" + dino
	holder := &spriteHolder{
		Avoid:     animations.LoadImage(path + "/" + "avoid.png"),
		Bite:      animations.LoadImage(path + "/" + "bite.png"),
		Dash:      animations.LoadImage(path + "/" + "dash.png"),
		Dead:      animations.LoadImage(path + "/" + "dead.png"),
		Hurt:      animations.LoadImage(path + "/" + "hurt.png"),
		Idle:      animations.LoadImage(path + "/" + "idle.png"),
		Jump:      animations.LoadImage(path + "/" + "jump.png"),
		Kick:      animations.LoadImage(path + "/" + "kick.png"),
		Move:      animations.LoadImage(path + "/" + "move.png"),
		Scan:      animations.LoadImage(path + "/" + "scan.png"),
		GhostIdle: animations.LoadImage(path + "/" + "ghost_idle.png"),
		GhostMove: animations.LoadImage(path + "/" + "ghost_move.png"),
	}
	holdersMap[dino] = holder
}
