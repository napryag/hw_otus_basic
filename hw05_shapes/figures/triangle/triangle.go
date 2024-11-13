package triangle

import "fmt"

type Triangle struct {
	Hight float64
	Base  float64
}

func (t *Triangle) NewTriangle(Hight float64, Base float64) {
	t.Hight = Hight
	t.Base = Base
}

func (t Triangle) CalculateArea() float64 {
	return (t.Base * t.Hight) / 2
}

func (t Triangle) Triangle() {
	fmt.Printf("Треугольник: высота %g, основание: %g\n", t.Hight, t.Base)
}
