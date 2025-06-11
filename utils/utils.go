package utils

import "math/rand"

const Pi = 3.1415926535897932385

func DegreesToRadians(degress float64) float64 {
	return degress * Pi / 180.0
}

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomFloat64InInterval(min, max float64) float64 {
	return min + (max-min)*RandomFloat64()
}
