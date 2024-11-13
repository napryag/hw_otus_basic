package circle

import "fmt"

const Pi float64 = 3.1415

type Circle struct {
	Radius float64
}

func (c *Circle) NewCircle(Radius float64) {
	c.Radius = Radius
}

func (c Circle) CalculateArea() float64 {
	return Pi * c.Radius * c.Radius
}

func (c Circle) Circle() {
	fmt.Printf("Круг: радиус %g\n", c.Radius)
}
