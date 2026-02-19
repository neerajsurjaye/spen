package model

type Circle struct {
	Pos    Position2D
	radius float64
}

func NewCircle(x float64, y float64, r float64) Circle {
	circle := Circle{
		Pos:    Position2D{x, y},
		radius: r,
	}

	return circle
}