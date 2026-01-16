package smath

import "math"

const Epsilon float64 = 1e-9

func EpsilonEqual(a float64, b float64) bool {
	return math.Abs(b - a) < Epsilon;
}

func LERP(a float64, b float64, t float64) float64{
	return a + (t * (b - a))
}

