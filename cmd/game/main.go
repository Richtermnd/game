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
	p2 := player.NewPlayer("cole", 200, 100)
	p3 := player.NewPlayer("kira", 300, 100)
	g.AddDrawer(p1, p2, p3)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
