package force

import "github.com/neerajsurjaye/spen/internal/smath"

type Gravity struct {
	Acceleration smath.Vec2
}

func NewGravity() Gravity {
	return Gravity{Acceleration: smath.Vec2{X: 0, Y: GRAVITATIONAL_ACCELERATION}}
}

func (g *Gravity) GetForce(mass float32) smath.Vec2 {
	return g.Acceleration.Multiply(mass)
}
