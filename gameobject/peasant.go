package gameobject

import (
	"github.com/Polifev/cereal/math"
)

const (
	Height = 16
	Width  = 16
)

var Dimensions = math.NewVectorFromInt(Width, Height)

type Peasant struct {
	position    math.Vector
	currentTask PeasantTask
}

func (p *Peasant) Contains(point math.Vector) bool {
	r := math.NewRectFromCenter(p.position, Dimensions)
	return r.Contains(point)
}

func NewPeasant(position math.Vector) *Peasant {
	return &Peasant{
		position:    position,
		currentTask: nil,
	}
}

func (p *Peasant) Position() math.Vector {
	return p.position
}

func (p *Peasant) Update() {
	if p.currentTask != nil {
		p.currentTask.Update(p)

		// TODO: pop instead
		if p.currentTask.Done() {
			p.currentTask = nil
		}
	}
}

func (p *Peasant) PushTask(task PeasantTask) {
	// TODO: push to queue instead
	p.currentTask = task
}

func (p *Peasant) Move(position math.Vector) {
	p.position = position
}
