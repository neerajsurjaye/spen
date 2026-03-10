package model

import (
	"github.com/neerajsurjaye/spen/internal/enums"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Circle struct {
	Transform Transform
	Color     Color
	AABB      AABB
	Body      *Body
}

func GetCircle(x float32, y float32, radius float32, rotation float32, col Color) *Circle {
	c := &Circle{}
	c.SetTransform(x, y, radius, rotation)
	c.SetColor(col)

	return c
}

func (c *Circle) GetColor() *Color{
	return &c.Color
}

func (c *Circle) SetColor(col Color) {
	c.Color = col
}

func (c *Circle) GetTransform() *Transform{
	return &c.Transform
}

func (c *Circle) SetTransform(x float32, y float32, radius float32, rot float32) {
	c.Transform = GetTransform(x, y, radius,radius, rot)
	c.Transform.Position.X = x
	c.Transform.Position.Y = y
	c.Transform.Radians = rot
	c.Transform.Scale = smath.Vec2{X: radius, Y: radius}
	c.updateAABB()
}

func (c *Circle) SetPosition(pos *smath.Vec2){
	c.Transform.Position = *pos
	c.updateAABB()
}

func (c *Circle) DeltaTransform(dx , dy, dsX ,dsY, dr float32){
	c.DeltaCircleTransform(dx, dy, dsX)
}

func (c *Circle) DeltaCircleTransform(dx float32, dy float32,  dr float32) {
	c.Transform.DeltaLocation(dx, dy, dr, dr, 0)
	c.updateAABB()
}

func (c *Circle) GetAABB() *AABB{
	return &c.AABB
}

func (c *Circle) updateAABB() {
	c.AABB.MinX = c.Transform.Position.X - c.Transform.Scale.X
	c.AABB.MaxX = c.Transform.Position.X + c.Transform.Scale.X
	c.AABB.MinY = c.Transform.Position.Y - c.Transform.Scale.Y
	c.AABB.MaxY = c.Transform.Position.Y + c.Transform.Scale.Y
}


func (c *Circle) SetBody(velX float32, velY float32, mass float32){
	c.Body = GetBody(velX, velY, mass)
}

func (c *Circle) GetBody() *Body{
	return c.Body
}

func (c *Circle) IsStatic() bool{
	return c.Body == nil
}

func (c *Circle) GetColliderType() enums.ColliderType{
	return enums.ColliderCircle
}

func (c *Circle) GetRadius() float32{
	return c.Transform.Scale.X
}