package render

import (
	"github.com/Polifev/cereal/gameobject"
	"github.com/Polifev/cereal/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PeasantRenderer struct {
	sprite *ebiten.Image
	offset math.Vector
}

func NewPeasantRenderer() (*PeasantRenderer, error) {
	image, _, err := ebitenutil.NewImageFromFile("resources/peasant.png")
	if err != nil {
		return nil, err
	}

	offset := math.NewVectorFromInt(image.Bounds().Dx()/2, image.Bounds().Dy()/2)
	return &PeasantRenderer{
		sprite: image,
		offset: offset,
	}, nil
}

func (pr *PeasantRenderer) Draw(peasant *gameobject.Peasant, screen *ebiten.Image) {
	geo := ebiten.GeoM{}
	p := peasant.Position().Sub(pr.offset)
	geo.Translate(p.X(), p.Y())
	screen.DrawImage(pr.sprite, &ebiten.DrawImageOptions{GeoM: geo})
}
