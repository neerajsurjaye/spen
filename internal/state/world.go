package state

import "github.com/neerajsurjaye/spen/internal/model"

type World struct {
	Circles []model.Circle
}

func (w *World) AddCircle(c model.Circle){
	w.Circles = append(w.Circles, c)
}