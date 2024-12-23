package main

import (
	"errors"
	"fmt"
)

func main() {
	h, w, err := scanValues(9, 4)
	if err != nil {
		fmt.Println("Не удалось создать доску по причине:", err)
	}
	fmt.Print(boardMaker(h, w))
}

func boardMaker(h, w int) [][]string {
	bb := make([][]string, h)
	for i := 0; i < h; i++ {
		bb[i] = make([]string, w)
		for j := 0; j < w; j++ {
			bb[i][j] = (boardWriter(i + j))
		}
	}
	return bb
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
