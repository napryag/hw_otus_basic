package main

import (
	"fmt"

	"github.com/napryag/hw_otus_basic/hw04_struct_comparator/types"
)

func main() {
	book1 := types.NewBook(1, "Зелёная Миля", "Стивен КИНГ", 1996, 384, 4.7)
	book2 := types.NewBook(2, "Крёстный отец", "Марио Пьюзо", 1969, 544, 4.6)

	comp1 := types.NewComparator(types.YEAR)
	comp2 := types.NewComparator(types.SIZE)
	comp3 := types.NewComparator(types.RATE)

	fmt.Printf("Книга %s выпущена раньше, чем книга %s: ", book1.Title(), book2.Title())
	fmt.Println(comp1.Compare(book1, book2))
	fmt.Printf("Книга %s имеет больше страниц, чем книга %s: ", book1.Title(), book2.Title())
	fmt.Println(comp2.Compare(book1, book2))
	fmt.Printf("Книга %s имеет больший рейтинг, чем книга %s: ", book1.Title(), book2.Title())
	fmt.Println(comp3.Compare(book1, book2))
}
