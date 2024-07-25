package field

import (
	"image/color"
	"log"
	"os"

	"github.com/Richtermnd/game/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	VOID = iota
	EMPTY
	OBSTACLE
)
const (
	tileSize = config.TileSize
	width    = config.FieldWidth
	heigth   = config.FieldHeight
)

type Field struct {
	field []byte //
}

func New(mapFilename string) *Field {
	field, err := os.ReadFile(mapFilename)
	if err != nil {
		log.Fatal(err)
	}
	if len(field) != width*heigth {
		log.Fatal("Wrong map size")
	}

	return &Field{
		field: field,
	}
}

func (f *Field) Draw(screen *ebiten.Image) {
	for i, tile := range f.field {
		x, y := i%width, i/width
		vector.DrawFilledRect(screen,
			float32(x*tileSize), float32(y*tileSize),
			float32(tileSize), float32(tileSize),
			getTileColor(tile), true)
	}
}

func (f *Field) Layout() int {
	return 0
}

func getTileColor(tile byte) color.Color {
	switch tile {
	case VOID:
		return color.Black
	case EMPTY:
		return color.RGBA{128, 128, 128, 255}
	case OBSTACLE:
		return color.RGBA{64, 64, 64, 255}
	default:
		return color.Black
	}
}
