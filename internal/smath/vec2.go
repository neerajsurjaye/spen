package smath

import (
	"math"
)

type Vec2 struct {
	X float64
	Y float64
}

func (a Vec2) Add(b Vec2) Vec2 {

	var result Vec2 = Vec2{a.X + b.X, a.Y + b.Y}
	return result
}

func (a Vec2) Multiply(magnitude float64) Vec2 {
	return Vec2{a.X * magnitude, a.Y * magnitude}
}

func (a Vec2) Divide(magnitude float64) Vec2 {
	if(math.Abs(magnitude) < Epsilon){
		return Vec2{0, 0}
	}

	return Vec2{a.X / magnitude, a.Y / magnitude}
}

func (a Vec2) Subtract(b Vec2) Vec2 {
	return a.Add(b.Multiply(-1))
}

func (a Vec2) Magnitude() float64 {

	var componentSquared float64 =  a.X * a.X + a.Y * a.Y
	return math.Sqrt(componentSquared);
}

func (a Vec2) Normalize() Vec2{
	aMagntude := a.Magnitude()

	if(aMagntude < Epsilon){
		return Vec2{0 , 0}
	}

	return a.Divide(aMagntude)
}

func (a Vec2) DistanceTo(point2 Vec2) float64{
	return (a.Subtract(point2)).Magnitude()
}

func (a Vec2) Dot(b Vec2) float64{
	return (a.X * b.X) + (a.Y * b.Y)
}

func (a Vec2) Angle(b Vec2) float64{
	
	dotProd := a.Dot(b)
	magnitude := (a.Magnitude() * b.Magnitude())

	if magnitude < Epsilon{
		return 0
	}

	cosine := dotProd / magnitude

	if cosine > 1{
		cosine = 1
	}else if (cosine < -1){
		cosine = -1
	}

	return math.Acos(cosine)
}

func (a Vec2) ProjectOn(b Vec2) Vec2{

	normalizedB := b.Normalize()

	aMagnitudeOnB := a.Dot(normalizedB)
	return normalizedB.Multiply(aMagnitudeOnB)
}

func (a Vec2) RejectOn(b Vec2) Vec2{
	return a.Subtract(a.ProjectOn(b));
}

func (a Vec2) ReflectOn(n Vec2) Vec2{
	return a.RejectOn(n).Subtract(a.ProjectOn(n))
}

func (a Vec2) DecomposeOn(b Vec2) [2]Vec2{
	return [2]Vec2{a.ProjectOn(b), a.RejectOn(b)}
}

func LERPV(a Vec2, b Vec2, t float64) Vec2{	
	return a.Add(b.Subtract(a).Multiply(t))
}