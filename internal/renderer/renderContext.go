package renderer

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/neerajsurjaye/spen/internal/model"
)

type RenderContext struct {
	OrthoProjection mgl32.Mat4
	ViewMat mgl32.Mat4
	Camera model.Camera
	DebugDraw bool
}