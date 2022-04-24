package figures

import "math"

type Circle struct {
	Radius float32
}

func (circle *Circle) area() float32 {
	return math.Pi * (circle.Radius * circle.Radius)
}

func (circle *Circle) perimeter() float32 {
	return 2 * math.Pi * circle.Radius
}
