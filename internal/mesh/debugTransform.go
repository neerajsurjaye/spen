package mesh

import (
	"math"

	"github.com/neerajsurjaye/spen/internal/model"
)

func DebugTransform(debugLines *DebugLines, t *model.Transform){

	xColor := model.GetColor(0.95, 0.25, 0.25, 0)
	yColor := model.GetColor(0.25, 0.95, 0.25, 0)


	//origin 
	ox, oy := transformPoint(0, 0 , t)

	//x axis
	xX, xY := transformPoint(1, 0, t)

	//y axis
	yX, yY := transformPoint(0, 1, t)

	debugLines.AddLine(ox, oy, xX, xY, xColor)
	debugLines.AddLine(ox, oy, yX, yY, yColor)
}

func transformPoint(x float32 , y float32, t *model.Transform) (float32, float32){
	cos := float32(math.Cos(float64(t.Radians)))
	sin := float32(math.Sin(float64(t.Radians)))

	//Just rotation matrix multiplication code
	rx := x * cos - y * sin
	ry := x * sin + y * cos

	rx *= t.Scale
	ry *= t.Scale

	rx += t.Position.X
	ry += t.Position.Y

	return rx, ry
}