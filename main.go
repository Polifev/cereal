package main

import (
	"fmt"
	"github.com/Polifev/cereal/gameobject"
	"github.com/Polifev/cereal/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type CerealGame struct {
	peasantRenderer *render.PeasantRenderer
	peasants        []*gameobject.Peasant
}

func newCerealGame() *CerealGame {
	game := CerealGame{}
	for i := 0; i < 10; i++ {
		game.peasants = append(game.peasants, gameobject.NewPeasant(100, float64(i*20)))
	}

	peasantRenderer, err := render.NewPeasantRenderer()
	if err != nil {
		log.Fatal(err)
	}
	game.peasantRenderer = peasantRenderer

	return &game
}

func (g *CerealGame) Update() error {
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
