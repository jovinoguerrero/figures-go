package figures

type Square struct {
	Width  float32
	Height float32
}

func (square *Square) area() float32 {
	return square.Width * square.Height
}

func (square *Square) perimeter() float32 {
	return 2*square.Width + 2*square.Height
}
