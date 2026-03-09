package model

import (
	"math"

	"github.com/neerajsurjaye/spen/internal/enums"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Wall struct {
	Color     Color
	Transform Transform
	AABB      AABB
	Body      *Body

}

func GetWall(x float32, y float32, scaleX, scaleY float32, rotation float32, col Color) *Wall {
	c := &Wall{}
	c.SetTransform(x, y, scaleX, scaleY, rotation)
	c.SetColor(col)

	return c
}

func (c *Wall) GetColor() *Color{
	return &c.Color
}

func (c *Wall) SetColor(col Color) {
	c.Color = col
}

func (c *Wall) GetTransform() *Transform{
	return &c.Transform
}

func (c *Wall) SetTransform(x float32, y float32, scaleX float32, scaleY float32, rot float32) {
	c.Transform = GetTransform(x, y, scaleX, scaleY, rot)
	c.Transform.Position.X = x
	c.Transform.Position.Y = y
	c.Transform.Radians = rot
	c.Transform.Scale = smath.Vec2{X: scaleX, Y: scaleY}
	c.updateAABB()
}

func (c *Wall) SetPosition(pos *smath.Vec2) {
	c.Transform.Position = *pos
	c.updateAABB()
}

func (c *Wall) DeltaTransform(dx, dy, dsX, dsY, dr float32) {
	c.Transform.DeltaLocation(dx, dy, dsX, dsY, dr)
	c.updateAABB()
}

func (c *Wall) updateAABB() {
	c.AABB.MinX = c.Transform.Position.X - c.Transform.Scale.X
	c.AABB.MaxX = c.Transform.Position.X + c.Transform.Scale.X
	c.AABB.MinY = c.Transform.Position.Y - c.Transform.Scale.Y
	c.AABB.MaxY = c.Transform.Position.Y + c.Transform.Scale.Y

	//Half width and half height
	hx := c.Transform.Scale.X
	hy := c.Transform.Scale.Y

	px := c.Transform.Position.X
	py := c.Transform.Position.Y

	cos := float32(math.Abs(math.Cos(float64(c.Transform.Radians))))
	sin := float32(math.Abs(math.Sin(float64(c.Transform.Radians))))

	newHx := cos * hx + sin * hy
	newHy := sin * hx + cos * hy

	c.AABB.MinX = px - newHx
	c.AABB.MaxX = px + newHx
	c.AABB.MinY = py - newHy
	c.AABB.MaxY = py + newHy
}


func (c *Wall) GetAABB() *AABB{
	return &c.AABB
}

func (c *Wall) GetBody() *Body{
	return c.Body
}

func (c *Wall) SetBody(velX float32, velY float32, mass float32){
	c.Body = GetBody(velX, velY, mass)
}

func (c *Wall) IsStatic() bool{
	return c.Body == nil
}

func (c *Wall) Type() enums.ColliderType{
	return enums.ColliderWall
}