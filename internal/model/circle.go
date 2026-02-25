package model

import (
	"github.com/neerajsurjaye/spen/internal/physics"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Circle struct {
	Transform Transform
	Color     Color
	AABB      AABB
	Body      physics.Body
}

func GetCircle(x float32, y float32, radius float32, rotation float32, col Color) *Circle {
	c := &Circle{}
	c.SetTransform(x, y, radius, rotation)
	c.SetColor(col)

	return c
}

func (c *Circle) SetColor(col Color) {
	c.Color = col
}

func (c *Circle) SetTransform(x float32, y float32, radius float32, rot float32) {
	c.Transform = GetTransform(x, y, radius, rot)
	c.Transform.Position.X = x
	c.Transform.Position.Y = y
	c.Transform.Radians = rot
	c.Transform.Scale = radius
	c.updateAABB()
}

func (c *Circle) SetPosition(pos *smath.Vec2){
	c.Transform.Position = *pos
	c.updateAABB()
}

func (c *Circle) DeltaTransform(dx float32, dy float32, dr float32) {
	c.Transform.DeltaLocation(dx, dy, dr, 0)
	c.updateAABB()
}

func (c *Circle) updateAABB() {
	c.AABB.MinX = c.Transform.Position.X - c.Transform.Scale
	c.AABB.MaxX = c.Transform.Position.X + c.Transform.Scale
	c.AABB.MinY = c.Transform.Position.Y - c.Transform.Scale
	c.AABB.MaxY = c.Transform.Position.Y + c.Transform.Scale
}


func (c *Circle) SetBody(velX float32, velY float32, mass float32){
	c.Body = physics.GetBody(velX, velY, mass)
}

func (c *Circle) GetBody() *physics.Body{
	return &c.Body
}