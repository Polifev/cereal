package gameobject

type Peasant struct {
	x float64
	y float64
}

func NewPeasant(x, y float64) *Peasant {
	return &Peasant{
		x: x,
		y: y,
	}
}

func (p *Peasant) X() float64 {
	return p.x
}

func (p *Peasant) Y() float64 {
	return p.y
}

func (p *Peasant) Position() (float64, float64) {
	return p.x, p.y
}

func (p *Peasant) Update() {
	p.x += 1.0
}
