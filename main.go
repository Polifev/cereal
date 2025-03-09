package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type CerealGame struct{}

func (g *CerealGame) Update() error {
	return nil
}

func (g *CerealGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cereal Game (%f)", ebiten.ActualFPS()), 0, 0)
}

func (g *CerealGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 300
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Cereal Game")
	err := ebiten.RunGame(&CerealGame{})
	if err != nil {
		log.Fatal(err)
	}
}
