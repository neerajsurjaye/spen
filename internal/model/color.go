package model

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func GetColor(r, g, b, a float32) *Color {
	return &Color{R: r, G: g, B: b, A: a}
}