package rectangle

import "fmt"

type Rectangle struct {
	FirstSide  float64
	SecondSide float64
}

func (r *Rectangle) NewRectangle(FirstSide float64, SecondSide float64) {
	r.FirstSide = FirstSide
	r.SecondSide = SecondSide
}

func (r Rectangle) CalculateArea() float64 {
	return r.FirstSide * r.SecondSide
}

func (r Rectangle) Rectangle() {
	fmt.Printf("Прямоугольник: первая сторона %g, вторая сторона: %g\n", r.FirstSide, r.SecondSide)
}
