package gameobject

import "github.com/Polifev/cereal/math"

type Tree struct {
	position math.Vector
}

func (t *Tree) Position() math.Vector {
	return t.position
}

func NewTree(position math.Vector) *Tree {
	return &Tree{position: position}
}
