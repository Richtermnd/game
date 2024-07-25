package utils

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Debugf(screen *ebiten.Image, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	ebitenutil.DebugPrint(screen, msg)
}
