package main

import (
	"fmt"

	"github.com/napryag/hw_otus_basic/hw04_struct_comparator/types"
)

const (
	year = iota + 1
	size
	rate
)

func main() {
	var book1, book2 types.Book
	book1.SetBook(1, "Зелёная Миля", "Стивен КИНГ", 1996, 384, 4.7)
	book2.SetBook(2, "Крёстный отец", "Марио Пьюзо", 1969, 544, 4.6)

	fmt.Println(book1.Book())
	fmt.Println(book2.Book())

	comp := types.Comparator{}
	comp.Compare(year, book1, book2)
	comp.Compare(size, book1, book2)
	comp.Compare(rate, book1, book2)
}
