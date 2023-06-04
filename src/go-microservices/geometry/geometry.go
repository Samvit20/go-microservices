package geometry

import "math"

//Area of rectangle
func Area(length, breadth float64) float64 {
	return length * breadth
}

//Diagonal of rectangle
func Diagonal(length, breadth float64) float64 {
	return math.Sqrt((length * length) + (breadth * breadth))
}
