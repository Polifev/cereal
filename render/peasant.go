package render

import (
	"github.com/Polifev/cereal/gameobject"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PeasantRenderer struct {
	sprite *ebiten.Image
}

func NewPeasantRenderer() (*PeasantRenderer, error) {
	image, _, err := ebitenutil.NewImageFromFile("resources/peasant.png")
	if err != nil {
		return nil, err
	}
	return &PeasantRenderer{sprite: image}, nil
}

func (pr *PeasantRenderer) Draw(peasant *gameobject.Peasant, screen *ebiten.Image) {
	geo := ebiten.GeoM{}
	geo.Translate(peasant.X(), peasant.Y())
	screen.DrawImage(pr.sprite, &ebiten.DrawImageOptions{GeoM: geo})
}
