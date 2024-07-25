package main

import (
	"log"

	"github.com/Richtermnd/game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	log.Default().SetPrefix("")
	g := game.New()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
