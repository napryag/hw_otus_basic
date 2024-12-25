package triangle

import "fmt"

type Triangle struct {
	Hight float64
	Base  float64
}

func (t *Triangle) NewTriangle(hight float64, base float64) {
	t.Hight = hight
	t.Base = base
}

func (t Triangle) CalculateArea() float64 {
	return (t.Base * t.Hight) / 2
}

func (t Triangle) Triangle() {
	fmt.Printf("Треугольник: высота %g, основание: %g\n", t.Hight, t.Base)
}
