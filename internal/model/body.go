package model

import (
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Body struct {
	Velocity   smath.Vec2
	Mass       float32
	Force      smath.Vec2
	Restituion float32
}

func GetBody(velX float32, velY float32, mass float32, restituion float32) *Body {
	return &Body{
		Velocity:   smath.Vec2{X: velX, Y: velY},
		Mass:       mass,
		Restituion: restituion,
	}
}

func (b *Body) AddForce(force smath.Vec2) {
	b.Force.AddSelf(force)
}

func (b *Body) SetVelocity(vel smath.Vec2) {
	b.Velocity = vel
}

func (b *Body) IntegrateVelocity(dt float32) {
	acceleration := b.Force.Multiply(1.0 / b.Mass)

	/*
		F = m * a
		a = F / m
	*/
	b.Velocity = b.Velocity.Add(acceleration.Multiply(dt))

	b.Force = smath.Vec2{}
}

func (b *Body) GetInvMass() float32 {
	if b.Mass != 0 {
		return 1 / b.Mass
	}
	return 0
}
