package state

import (
	"github.com/neerajsurjaye/spen/internal/mesh"
)

type World struct {
	Circles []mesh.CircleInfo
}

func (w *World) AddCircle(c mesh.CircleInfo){
	w.Circles = append(w.Circles, c)
}