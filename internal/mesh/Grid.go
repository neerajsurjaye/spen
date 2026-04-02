package mesh

import (
	"math"

	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type Grid struct {
	spacing      float32
	visibleRange float32
}

func (g *Grid) Init(spacing float32, visibleRange float32) {
	g.spacing = spacing
	g.visibleRange = visibleRange
}

func (g *Grid) Debug(rc *renderer.RenderContext, debugLines *DebugLines) {

	camera := rc.Camera

	startX := float32(math.Floor(float64(camera.X/g.spacing))) * float32(g.spacing)
	startY := float32(math.Floor(float64(camera.Y/g.spacing))) * float32(g.spacing)

	gridColor := model.GetColor(0.8, 0.8, 0.8, 1)
	xAxisCol := model.GetColor(1, 0, 0, 1)
	yAxisCol := model.GetColor(0, 1, 0, 1)

	for x := startX - g.visibleRange; x <= startX+g.visibleRange; x += g.spacing {
		if smath.EpsilonEqualF32(x, 0) {
			debugLines.AddLine(x, camera.Y-g.visibleRange, x, camera.Y+g.visibleRange, yAxisCol)
			continue
		}
		debugLines.AddLine(x, camera.Y-g.visibleRange, x, camera.Y+g.visibleRange, gridColor)
	}

	for y := startY - g.visibleRange; y <= startY+g.visibleRange; y += g.spacing {
		if smath.EpsilonEqualF32(y, 0) {
			debugLines.AddLine(camera.X-g.visibleRange, y, camera.X+g.visibleRange, y, xAxisCol)
			continue
		}
		debugLines.AddLine(camera.X-g.visibleRange, y, camera.X+g.visibleRange, y, gridColor)
	}
}

func (g *Grid) GetSpacing() float32 {
	return g.spacing
}

func (g *Grid) SetSpacing(spacing float32) {
	g.spacing = spacing
}
