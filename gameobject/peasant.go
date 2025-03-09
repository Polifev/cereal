package gameobject

import "github.com/Polifev/cereal/math"

type Peasant struct {
	position    math.Vector
	currentTask PeasantTask
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
