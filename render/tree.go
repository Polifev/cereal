package render

import (
	"github.com/Polifev/cereal/gameobject"
	"github.com/Polifev/cereal/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TreeRenderer struct {
	trunkSprite  *ebiten.Image
	leavesSprite *ebiten.Image
	trunkOffset  math.Vector
	leavesOffset math.Vector
}

func NewTreeRenderer() (*TreeRenderer, error) {
	trunk, _, err := ebitenutil.NewImageFromFile("resources/tree_trunk.png")
	if err != nil {
		return nil, err
	}

	leaves, _, err := ebitenutil.NewImageFromFile("resources/tree_leaves.png")
	if err != nil {
		return nil, err
	}

	trunkOffset := math.NewVectorFromInt(trunk.Bounds().Dx()/2, trunk.Bounds().Dy()/2)
	leavesOffset := trunkOffset.Add(math.NewVectorFromInt(0, leaves.Bounds().Dy()))

	return &TreeRenderer{
		trunkSprite:  trunk,
		leavesSprite: leaves,
		trunkOffset:  trunkOffset,
		leavesOffset: leavesOffset,
	}, nil
}

func (tr *TreeRenderer) Draw(tree *gameobject.Tree, screen *ebiten.Image) {
	geo := ebiten.GeoM{}
	trunkPos := tree.Position().Sub(tr.trunkOffset)
	leavesPos := tree.Position().Sub(tr.leavesOffset)
	geo.Translate(trunkPos.X(), trunkPos.Y())
	screen.DrawImage(tr.trunkSprite, &ebiten.DrawImageOptions{GeoM: geo})
	geo.Reset()
	geo.Translate(leavesPos.X(), leavesPos.Y())
	screen.DrawImage(tr.leavesSprite, &ebiten.DrawImageOptions{GeoM: geo})
}
