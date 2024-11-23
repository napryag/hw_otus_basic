package main

import (
	"errors"
	"fmt"
)

func main() {
	h, w, err := scanValues()
	if err != nil {
		fmt.Println("Не удалось создать доску по причине:", err)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(boardWriter(i + j))
		}
		fmt.Println()
	}
}

func boardWriter(n int) string {
	if n%2 == 0 {
		return "#"
	}
	return " "
}

func scanValues() (int, int, error) {
	var hight, width int
	fmt.Println("Укажите размер шахматной доски. Введите высоту доски: ")
	fmt.Scan(&hight)
	fmt.Println("Введите ширину доски: ")
	fmt.Scan(&width)
	if hight <= 0 || width <= 0 {
		err := errors.New("значение высоты и ширины должно быть положительным целым числом")
		return hight, width, err
	}
	return hight, width, nil
}
