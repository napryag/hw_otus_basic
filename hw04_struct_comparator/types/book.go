package types

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetBook(id int, title string, author string, year int, size int, rate float32) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

func (b Book) Book() string {
	return fmt.Sprintf(
		"ID: %d; Title: %s; Author: %s; Year: %d; Size: %d; Rate: %g",
		b.id,
		b.title,
		b.author,
		b.year,
		b.size,
		b.rate,
	)
}

type Comparator struct{}

func (c Comparator) Compare(k int, b1, b2 Book) {
	switch k {
	case 1:
		fmt.Printf("Книга %s издана позже, чем книга %s: ", b1.title, b2.title)
		fmt.Println(b1.year > b2.year)
	case 2:
		fmt.Printf("Книга %s имеет больше страниц, чем книга %s: ", b1.title, b2.title)
		fmt.Println(b1.size > b2.size)
	case 3:
		fmt.Printf("Книга %s имеет больший рейтинг, чем книга %s: ", b1.title, b2.title)
		fmt.Println(b1.rate > b2.rate)
	}
}
