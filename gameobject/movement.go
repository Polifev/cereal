package gameobject

import "github.com/Polifev/cereal/nav"

type MovementTask struct {
	path nav.Path
}

func NewMovementTask(path nav.Path) MovementTask {
	return MovementTask{path: path}
}

func (m *MovementTask) Update(p *Peasant) {
	if m.path.IsDone() {
		// does nothing
		// should not happen
		return
	}

	// TODO: export anywhere else
	peasantSpeed := 2.0

	diff := m.path.Current().Sub(p.Position())
	actualSpeed := min(peasantSpeed, diff.Length()) // avoid passing by the target
	// TODO: use remaining speed for next path node
	movement := diff.Normalized().Times(actualSpeed)
	newPosition := p.Position().Add(movement)
	p.Move(newPosition)
	m.path.Update(newPosition)
}

func (m *MovementTask) Done() bool {
	return m.path.IsDone()
}
