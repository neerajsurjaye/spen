package force

import "github.com/neerajsurjaye/spen/internal/smath"

type Gravity struct {
	Acceleration smath.Vec2
}

func NewGravity(gravAcc *smath.Vec2) Gravity {
	if gravAcc == nil {
		return Gravity{Acceleration: smath.Vec2{X: 0, Y: GRAVITATIONAL_ACCELERATION}}
	}
	return Gravity{Acceleration: *gravAcc}
}

func (g *Gravity) GetForce(mass float32) smath.Vec2 {
	return g.Acceleration.Multiply(mass)
}
