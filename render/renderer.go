package render

import "github.com/hajimehoshi/ebiten/v2"

type Renderer[T interface{}] interface {
	Draw(element *T, screen *ebiten.Image)
}
