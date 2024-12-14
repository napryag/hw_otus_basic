package main

import (
	"errors"
	"fmt"
)

func main() {
	h, w, err := scanValues(8, 8)
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

func scanValues(hight, width int) (int, int, error) {
	fmt.Printf("Укажите размер шахматной доски. Введите высоту доски: %d\n", hight)
	if hight <= 0 {
		err := errors.New("значение высоты должно быть положительным целым числом")
		return hight, width, err
	}
	fmt.Printf("Введите ширину доски: %d\n", width)

	if width <= 0 {
		err := errors.New("значение ширины должно быть положительным целым числом")
		return hight, width, err
	}
	return hight, width, nil
}
