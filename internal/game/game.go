package game

import (
	"cmp"
	"slices"

	"github.com/Richtermnd/game/internal/config"
	"github.com/Richtermnd/game/internal/field"
	"github.com/hajimehoshi/ebiten/v2"
)

type LayoutDrawer interface {
	Layout() int
	Draw(screen *ebiten.Image)
}

type Game struct {
	drawers []LayoutDrawer
}

func New() *Game {
	setup()
	g := &Game{}
	f := field.New("maps/map.gf")
	g.AddDrawer(f)
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, drawer := range g.drawers {
		drawer.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) AddDrawer(drawer LayoutDrawer) {
	g.drawers = append(g.drawers, drawer)
	slices.SortFunc(g.drawers, func(a, b LayoutDrawer) int {
		return cmp.Compare(a.Layout(), b.Layout())
	})
}

func setup() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Game")
}
