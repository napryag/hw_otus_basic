package main

import "fmt"

func makeChessboard(h int, w int) {
	even := []string{}

	for k := 0; k < (w / 2); k++ {
		even = append(even, "  #")
	}

	uneven := []string{}

	for j := 0; j < (w / 2); j++ {
		uneven = append(uneven, "#  ")
	}

	for i := 1; i <= h; i++ {
		if i%2 != 0 {
			fmt.Println(uneven)
		} else {
			fmt.Println(even)
		}
	}
}

func main() {
	var hight int
	var width int

	fmt.Println("Введите размер поля. \n", "Укажите высоту и ширину через пробел: ")

	fmt.Scanf("%d %d", &hight, &width)

	makeChessboard(hight, width)
}
