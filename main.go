package main

import (
	"fmt"
	"github.com/Polifev/cereal/gameobject"
	"github.com/Polifev/cereal/math"
	"github.com/Polifev/cereal/nav"
	"github.com/Polifev/cereal/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type CerealGame struct {
	navigation      nav.Navigation
	peasantRenderer *render.PeasantRenderer
	peasants        []*gameobject.Peasant
}

func newCerealGame() *CerealGame {
	game := CerealGame{}
	for i := 0; i < 10; i++ {
		peasant := gameobject.NewPeasant(math.NewVectorFromInt(60, i*20))
		game.peasants = append(game.peasants, peasant)
	}

	peasantRenderer, err := render.NewPeasantRenderer()
	if err != nil {
		log.Fatal(err)
	}
	game.peasantRenderer = peasantRenderer

	// Define navigation engine (will later use A-Star or similar)
	navigation := nav.Straight{}
	game.navigation = &navigation

	return &game
}

func (g *CerealGame) Update() error {
	// TODO: move this in an appropriate location
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// TODO: compute real-world coordinate using camera
		x, y := ebiten.CursorPosition()
		target := math.NewVectorFromInt(x, y)
		peasant := g.peasants[0]
		path := g.navigation.Compute(peasant.Position(), target)
		newTask := gameobject.NewMovementTask(path)
		peasant.PushTask(&newTask)
	}

	for i := 0; i < len(g.peasants); i++ {
		g.peasants[i].Update()
	}
	return nil
}

func (g *CerealGame) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.peasants); i++ {
		g.peasantRenderer.Draw(g.peasants[i], screen)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cereal Game (%f)", ebiten.ActualFPS()), 0, 0)
}

func (g *CerealGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 300
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Cereal Game")

	game := newCerealGame()
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
