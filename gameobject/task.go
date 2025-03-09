package gameobject

type PeasantTask interface {
	Update(p *Peasant)
	Done() bool
}
