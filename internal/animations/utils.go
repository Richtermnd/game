package animations

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
