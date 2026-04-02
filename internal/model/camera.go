package model

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	X    float32
	Y    float32
	Zoom float32
}

func (c *Camera) View() mgl32.Mat4 {

	translation := mgl32.Translate3D(-c.X, -c.Y, 0)
	zoom := mgl32.Scale3D(c.Zoom, c.Zoom, 1)

	return zoom.Mul4(translation)
}

func (c *Camera) MoveDelta(dx float32, dy float32) {
	c.X += dx
	c.Y += dy

	// fmt.Println(c)
}

func (c *Camera) ZoomDelta(dZoom float32) {
	c.Zoom += dZoom
}
