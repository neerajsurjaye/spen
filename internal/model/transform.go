package model

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Transform struct {
	Position smath.Vec2

	Scale float32

	//radians
	Radians float32
}

func GetTransform(x float32, y float32, scale float32, radians float32) Transform{
	return Transform{
		Position: smath.Vec2{X: x, Y : y},
		Scale: scale,
		Radians: radians,
	}
}

func (t *Transform) DeltaLocation(dx float32, dy float32, ds float32, dr float32) {
	t.Position.X = t.Position.X + dx
	t.Position.Y = t.Position.Y + dy
	t.Scale = t.Scale + ds
	t.Radians = t.Radians + dr
}

func (t *Transform) GetModel() mgl32.Mat4 {
	return t.getTranslation().Mul4(t.getRotation()).Mul4(t.getScale())
}

func (t *Transform) getTranslation() mgl32.Mat4 {
	return mgl32.Translate3D(float32(t.Position.X), float32(t.Position.Y), 0)
}

func (t *Transform) getScale() mgl32.Mat4 {
	return mgl32.Scale3D(float32(t.Scale), float32(t.Scale), 1)
}

func (t *Transform) getRotation() mgl32.Mat4 {
	return mgl32.HomogRotate3DZ(t.Radians)
}