package model

import "github.com/go-gl/mathgl/mgl32"

type Circle struct {
	X float32
	Y float32
	R float32

	Color Color
	AABB AABB
}

func GetCircle(x float32, y float32, r float32, col Color) *Circle{
	c := &Circle{}
	c.SetTransform(x, y , r)
	c.SetColor(col)

	return c
}

func (c *Circle) SetColor(col Color) {
	c.Color = col
}

func (c *Circle) SetTransform(x float32, y float32, r float32) {
	c.X = x
	c.Y = y
	c.R = r

	c.updateAABB()
}

func (c *Circle) updateAABB() {
	c.AABB.MinX = c.X - c.R
	c.AABB.MaxX = c.X + c.R
	c.AABB.MinY = c.Y - c.R
	c.AABB.MaxY = c.Y + c.R
}

func (c *Circle) GetModel() mgl32.Mat4 {
	return c.getTranslation().Mul4(c.getScale())
}

func (c *Circle) getTranslation() mgl32.Mat4 {
	return mgl32.Translate3D(float32(c.X), float32(c.Y), 0)
}

func (c *Circle) getScale() mgl32.Mat4 {
	return mgl32.Scale3D(float32(c.R), float32(c.R), 1)
}

