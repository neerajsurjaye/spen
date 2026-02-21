package state

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/model"
)

/*
Contains information related to the world running inside engine.
*/
type World struct {
	Camera model.Camera
	Circles []mesh.CircleInfo
	Background model.Color
}

var world *World = nil

func InitWorld(){
	if world == nil{
		world = &World{
			Circles: []mesh.CircleInfo{},
		}
	}
}

func WorldInstance() *World{
	return world
}

func (w *World) AddCircle(c mesh.CircleInfo){
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
	for circle := range(w.Circles){
		w.Circles[circle].Draw(w.GetOrthoProjection(), w.Camera.View())
	}

}