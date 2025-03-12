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
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type CerealGame struct {
	hoveredPeasant *gameobject.Peasant

	navigation      nav.Navigation
	peasantRenderer *render.PeasantRenderer
	treeRenderer    *render.TreeRenderer
	peasants        []*gameobject.Peasant
	trees           []*gameobject.Tree
}

func newCerealGame() *CerealGame {
	game := CerealGame{}
	for i := 0; i < 10; i++ {
		peasant := gameobject.NewPeasant(math.NewVectorFromInt(60, i*20))
		game.peasants = append(game.peasants, peasant)
	}

	for i := 0; i < 10; i++ {
		tree := gameobject.NewTree(math.NewVectorFromInt(i*40, 500-i*40))
		game.trees = append(game.trees, tree)
	}
	game.trees = append(game.trees, gameobject.NewTree(math.NewVectorFromInt(0, 0)))

	peasantRenderer, err := render.NewPeasantRenderer()
	if err != nil {
		log.Fatal(err)
	}
	game.peasantRenderer = peasantRenderer

	treeRenderer, err := render.NewTreeRenderer()
	if err != nil {
		log.Fatal(err)
	}
	game.treeRenderer = treeRenderer

	// Define navigation engine (will later use A-Star or similar)
	navigation := nav.Straight{}
	game.navigation = &navigation

	return &game
}

func (g *CerealGame) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	mousePos := math.NewVectorFromInt(mouseX, mouseY)

	g.hover(mousePos)

	// TODO: move this in an appropriate location
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// TODO: compute real-world coordinate using camera
		peasant := g.peasants[0]
		path := g.navigation.Compute(peasant.Position(), mousePos)
		newTask := gameobject.NewMovementTask(path)
		peasant.PushTask(&newTask)
	}

	for i := 0; i < len(g.peasants); i++ {
		g.peasants[i].Update()
	}
	return nil
}

func (g *CerealGame) hover(mousePos math.Vector) {
	for _, peasant := range g.peasants {
		if peasant.Contains(mousePos) {
			g.hoveredPeasant = peasant
			return
		}
	}
	g.hoveredPeasant = nil
}

func (g *CerealGame) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.peasants); i++ {
		peasant := g.peasants[i]
		pos := peasant.Position()
		if g.hoveredPeasant == g.peasants[i] {
			vector.DrawFilledCircle(screen, float32(pos.X()), float32(pos.Y()), 12, color.RGBA{255, 255, 255, 255}, false)
		}
		g.peasantRenderer.Draw(g.peasants[i], screen)
	}
	for i := 0; i < len(g.trees); i++ {
		g.treeRenderer.Draw(g.trees[i], screen)
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
