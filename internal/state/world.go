package state

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
)

/*
Contains information related to the world running inside engine.
*/
type World struct {
	Camera model.Camera
	Circles []*model.Circle
	CircleMesh mesh.CircleMesh
	DebugLines mesh.DebugLines
	Background model.Color
	Grid mesh.Grid
	DebugDraw bool
}

var world *World = nil

func InitWorld(){
	if world == nil{
		world = &World{}
	}
}

func WorldInstance() *World{
	return world
}

func (w *World) AddCircle(c *model.Circle){
	w.Circles = append(w.Circles, c)
}

func (w *World) SetBackground(r float32, g float32, b float32, a float32){
	w.Background = model.Color{R : r, G : g, B : b, A : a}
}

func (w *World) SetCamera(x float32, y float32){
	w.Camera.X = x
	w.Camera.Y = y
}

func (w *World) GetOrthoProjection() mgl32.Mat4{
	engine := EngineInstance()

	//-width/2 to width/2
	halfW := float32(engine.UiInfo.Width) / 2
	halfH := float32(engine.UiInfo.Height) / 2 

	return mgl32.Ortho(
		-halfW, halfW,
		-halfH, halfH,
		-1 , 1,
	)
}

func (w *World) Draw(){


	renderContext := &renderer.RenderContext{
		OrthoProjection: w.GetOrthoProjection(),
		ViewMat: w.Camera.View(),
		Camera: w.Camera,
		DebugDraw: w.DebugDraw,
	}

	for idx := range(w.Circles){
		w.CircleMesh.Draw(renderContext, w.Circles[idx])
	}

	if w.DebugDraw{
		var aabbDebugColor model.Color = model.Color{R: 1.0, G: 1.0, B: 0.2, A: 1} 

		w.Grid.Debug(renderContext, &w.DebugLines)
		for idx := range(w.Circles){
			mesh.DebugAABB(&w.DebugLines, &w.Circles[idx].AABB, &aabbDebugColor)
			mesh.DebugTransform(&w.DebugLines, &w.Circles[idx].Transform)
		}
		w.DebugLines.Draw(renderContext)
	}
	
}