package model

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	X float32
	Y float32
}

func (c *Camera) View() mgl32.Mat4 {
	//Translate world by inverse of camera location
	return mgl32.Translate3D(-c.X, -c.Y, 0)
}

func (c *Camera) MoveDelta(dx float32, dy float32) {
	c.X += dx
	c.Y += dy

	// fmt.Println(c)
}
