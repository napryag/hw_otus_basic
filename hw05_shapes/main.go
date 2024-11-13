package main

import (
	"errors"
	"fmt"

	"github.com/fixme_my_friend/hw05_shapes/figures"
	"github.com/fixme_my_friend/hw05_shapes/figures/circle"
	"github.com/fixme_my_friend/hw05_shapes/figures/rectangle"
	"github.com/fixme_my_friend/hw05_shapes/figures/square"
	"github.com/fixme_my_friend/hw05_shapes/figures/triangle"
)

func main() {
	var circle circle.Circle
	var rectangle rectangle.Rectangle
	var triangle triangle.Triangle
	var square square.Square

	circle.NewCircle(4)
	rectangle.NewRectangle(10, 5)
	triangle.NewTriangle(7, 12)
	square.NewSquare(12)

	a, err := calculateArea(circle)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	circle.Circle()
	fmt.Printf("Площадь: %g\n", a)
	b, err := calculateArea(rectangle)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	rectangle.Rectangle()
	fmt.Printf("Площадь: %g\n", b)
	c, err := calculateArea(triangle)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	triangle.Triangle()
	fmt.Printf("Площадь: %g\n", c)
	d, err := calculateArea(square)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	square.Square()
	fmt.Printf("Площадь: %g\n", d)
}

func calculateArea(a any) (float64, error) {
	if figure, ok := a.(figures.Shape); ok {
		b := figure.CalculateArea()
		return b, nil
	}
	err := errors.New("переданный объект не имплементирует интерфейс Shape")
	return 0, err
}
