package game

import (
	"cmp"
	"slices"

	"github.com/Richtermnd/game/internal/config"
	"github.com/Richtermnd/game/internal/field"
	"github.com/Richtermnd/game/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type LayoutDrawer interface {
	Layout() int
	Draw(screen *ebiten.Image)
}

type Updater interface {
	Update() error
}

type Game struct {
	drawers  []LayoutDrawer
	updaters []Updater
}

func New() *Game {
	setup()
	g := &Game{}
	f := field.New("maps/map.gf")
	g.AddDrawer(f)
	return g
}

func (g *Game) Update() error {
	for _, updater := range g.updaters {
		err := updater.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, drawer := range g.drawers {
		drawer.Draw(screen)
	}
	utils.DrawDebug(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) AddDrawer(drawer ...LayoutDrawer) {
	g.drawers = append(g.drawers, drawer...)
	slices.SortFunc(g.drawers, func(a, b LayoutDrawer) int {
		return cmp.Compare(a.Layout(), b.Layout())
	})
}
func (g *Game) AddUpdaters(updaters ...Updater) {
	g.updaters = append(g.updaters, updaters...)
}

func setup() {
	ebiten.SetTPS(30)
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Game")
}
