package smath

import "math"

const EpsilonF64 float64 = 1e-9
const EpsilonF32 float32 = 1e-3

func EpsilonEqualF64(a float64, b float64) bool {
	return math.Abs(b - a) < EpsilonF64;
}


func EpsilonEqualF32(a float32, b float32) bool {
	return float32(math.Abs(float64(b - a))) < EpsilonF32;
}

func LERP(a float64, b float64, t float64) float64{
	return a + (t * (b - a))
}

