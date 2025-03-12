package math

type Rect struct {
	position  Vector
	dimension Vector
}

func NewRectFromInt(x, y, w, h int) Rect {
	return Rect{
		position:  NewVectorFromInt(x, y),
		dimension: NewVectorFromInt(w, h),
	}
}

func NewRectFromFloat(x, y, w, h float64) Rect {
	return Rect{
		position:  NewVector(x, y),
		dimension: NewVector(w, h),
	}
}

func NewRect(position Vector, dimension Vector) Rect {
	return Rect{position, dimension}
}

func NewRectFromCenter(position Vector, dimension Vector) Rect {
	return Rect{
		position:  position.Sub(dimension.Times(0.5)),
		dimension: dimension,
	}
}

func (rect *Rect) Position() Vector {
	return rect.position
}

func (rect *Rect) Dimension() Vector {
	return rect.dimension
}

func (rect *Rect) Contains(vector Vector) bool {
	high := rect.position.Add(rect.dimension)

	withinX := vector.X() >= rect.position.X() && vector.X() <= high.X()
	withinY := vector.Y() >= rect.position.Y() && vector.Y() <= high.Y()

	return withinY && withinX
}
