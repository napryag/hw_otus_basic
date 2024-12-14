package square

import "fmt"

type Square struct {
	Side float64
}

func (s *Square) NewSquare(side float64) {
	s.Side = side
}

func (s Square) CalculatePerimeter() float64 {
	return s.Side * 4
}

func (s Square) Square() {
	fmt.Printf("Квадрат: сторона %g\n", s.Side)
}
