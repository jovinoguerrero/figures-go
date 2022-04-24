package figures

import "fmt"

type geometricFigure interface {
	area() float32
	perimeter() float32
}

func GetMetrics(geometricFigure geometricFigure) {
	fmt.Println("Metrics:", geometricFigure)
	fmt.Println("Area:", geometricFigure.area())
	fmt.Println("Perimeter:", geometricFigure.perimeter())
}