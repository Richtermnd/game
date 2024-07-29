package main

import (
	"log"

	"github.com/Richtermnd/game/internal/game"
	"github.com/Richtermnd/game/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	log.Default().SetPrefix("")
	g := game.New()
	p1 := player.NewPlayer("olaf", 100, 100)
	g.AddDrawer(p1)
	g.AddUpdaters(p1)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
