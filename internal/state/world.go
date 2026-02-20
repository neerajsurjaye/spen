package state

import (
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/model"
)

type World struct {
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

func (w *World) Draw(){
	for circle := range(w.Circles){
		w.Circles[circle].Draw()
	}

}