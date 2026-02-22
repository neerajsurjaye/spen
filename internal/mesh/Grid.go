package mesh

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/utils"
)

type Grid struct {
	spacing      float32
	visibleRange float32

	debugDraw DebugLines
}

func (g *Grid)Init(spacing float32, visibleRange float32) {
	g.spacing = spacing
	g.visibleRange = visibleRange
	g.debugDraw = g.debugDraw
	g.debugDraw.Init(utils.CreateProgram("shader/grid/grid.frag" , "shader/grid/grid.vert"))

}

func (g *Grid) Draw(camera model.Camera, projection mgl32.Mat4, view mgl32.Mat4) {

	startX := float32(math.Floor(float64(camera.X / g.spacing))) * float32(g.spacing)
	startY := float32(math.Floor(float64(camera.Y / g.spacing))) * float32(g.spacing)

	for x := startX - g.visibleRange; x <= startX + g.visibleRange; x += g.spacing{
		g.debugDraw.AddLine(x , camera.Y - g.visibleRange, x , camera.Y + g.visibleRange)	
	}

	for y := startY - g.visibleRange; y <= startY + g.visibleRange; y += g.spacing{
		g.debugDraw.AddLine(camera.X - g.visibleRange, y, camera.X + g.visibleRange, y)
	}

	g.debugDraw.Draw(projection, view)
}