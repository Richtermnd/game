package animations

import (
	"image/png"
	"log"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	holdersMap map[string]*SpriteHolder = map[string]*SpriteHolder{}
	mu         sync.Mutex
)

type SpriteHolder struct {
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

const spriteFolder = "sprites"

func init() {
	dinos, err := os.ReadDir("sprites")
	if err != nil {
		log.Fatal(err)
	}
	for _, dino := range dinos {
		uploadDino(dino.Name())
	}
}

func getSpriteHolder(name string) *SpriteHolder {
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
	holder := &SpriteHolder{
		Avoid:     loadImage(path + "/" + "avoid.png"),
		Bite:      loadImage(path + "/" + "bite.png"),
		Dash:      loadImage(path + "/" + "dash.png"),
		Dead:      loadImage(path + "/" + "dead.png"),
		Hurt:      loadImage(path + "/" + "hurt.png"),
		Idle:      loadImage(path + "/" + "idle.png"),
		Jump:      loadImage(path + "/" + "jump.png"),
		Kick:      loadImage(path + "/" + "kick.png"),
		Move:      loadImage(path + "/" + "move.png"),
		Scan:      loadImage(path + "/" + "scan.png"),
		GhostIdle: loadImage(path + "/" + "ghost_idle.png"),
		GhostMove: loadImage(path + "/" + "ghost_move.png"),
	}
	holdersMap[dino] = holder
}

func loadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal("animations.uploading.uploadImage:", err)
	}
	return ebiten.NewImageFromImage(img)
}
