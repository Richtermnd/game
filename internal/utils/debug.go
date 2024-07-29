package utils

import (
	"fmt"

	"github.com/Richtermnd/game/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var messages []string

func Debugf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	messages = append(messages, msg)
}

func DrawDebug(screen *ebiten.Image) {
	for i, msg := range messages {
		y := i * 10
		x := i / config.ScreenHeight * 200
		ebitenutil.DebugPrintAt(screen, msg, x, y)
	}
	messages = messages[:0] // reset
}
