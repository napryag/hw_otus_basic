package square

import "fmt"

type Square struct {
	Side float64
}

func (s *Square) NewSquare(Side float64) {
	s.Side = Side
}

func (s Square) CalculatePerimeter() float64 {
	return s.Side * s.Side * s.Side * s.Side
}

func (s Square) Square() {
	fmt.Printf("Квадрат: сторона %g\n", s.Side)
}
