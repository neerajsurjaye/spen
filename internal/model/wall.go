package model

import (
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Wall struct {
	Color     Color
	Transform Transform
	AABB      AABB
}

func GetWall(x float32, y float32, scaleX, scaleY float32, rotation float32, col Color) *Wall {
	c := &Wall{}
	c.SetTransform(x, y, scaleX, scaleY, rotation)
	c.SetColor(col)

	return c
}

func (c *Wall) SetColor(col Color) {
	c.Color = col
}

func (c *Wall) SetTransform(x float32, y float32, scaleX float32, scaleY float32, rot float32) {
	c.Transform = GetTransform(x, y, scaleX, scaleY, rot)
	c.Transform.Position.X = x
	c.Transform.Position.Y = y
	c.Transform.Radians = rot
	c.Transform.Scale = smath.Vec2{scaleX, scaleY}
	c.updateAABB()
}

func (c *Wall) SetPosition(pos *smath.Vec2) {
	c.Transform.Position = *pos
	c.updateAABB()
}

func (c *Wall) DeltaTransform(dx float32, dy float32, dsX, dsY, dr float32) {
	c.Transform.DeltaLocation(dx, dy, dsX, dsY, dr)
	c.updateAABB()
}

func (c *Wall) updateAABB() {
	c.AABB.MinX = c.Transform.Position.X - c.Transform.Scale.X
	c.AABB.MaxX = c.Transform.Position.X + c.Transform.Scale.X
	c.AABB.MinY = c.Transform.Position.Y - c.Transform.Scale.Y
	c.AABB.MaxY = c.Transform.Position.Y + c.Transform.Scale.Y
}
