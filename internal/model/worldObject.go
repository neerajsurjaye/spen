package model

import (
	"github.com/neerajsurjaye/spen/internal/physics"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type WorldObject interface {
	GetTransform() *Transform
	
	GetColor() *Color
	SetColor(col Color)
	
	SetPosition(pos *smath.Vec2)
	DeltaTransform(dx , dy, dsX ,dsY, dr float32)

	GetAABB() *AABB

	GetBody() *physics.Body

	IsStatic() bool
}