package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBook(t *testing.T) {
	testCases := []struct {
		name  string
		examp Book
		want  Book
		got   Book
	}{
		{
			name:  "1",
			examp: Book{144, "Горе от ума", "Грибоедов", 1825, 154, 4.6},
			want:  Book{144, "Горе от ума", "Грибоедов", 1825, 154, 4.6},
		},
		{
			name:  "2",
			examp: Book{12, "Евгений Онегин", "Пушкин", 1825, 215, 4.5},
			want:  Book{12, "Евгений Онегин", "Пушкин", 1825, 215, 4.5},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = *NewBook(tC.examp.id, tC.examp.title, tC.examp.author, tC.examp.year, tC.examp.size, tC.examp.rate)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_SetID(t *testing.T) {
	testCases := []struct {
		name string
		ID   int
		want Book
		got  Book
	}{
		{
			name: "1",
			ID:   13,
			want: Book{id: 13},
		},
		{
			name: "2",
			ID:   228,
			want: Book{id: 228},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetID(tC.ID)
			assert.Equal(t, tC.want.id, tC.got.id)
		})
	}
}

func Test_SetTitle(t *testing.T) {
	testCases := []struct {
		name  string
		title string
		want  Book
		got   Book
	}{
		{
			name:  "1",
			title: "Биология",
			want:  Book{title: "Биология"},
		},
		{
			name:  "2",
			title: "Физика",
			want:  Book{title: "Физика"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetTitle(tC.title)
			assert.Equal(t, tC.want.title, tC.got.title)
		})
	}
}

func Test_SetAuthor(t *testing.T) {
	testCases := []struct {
		name   string
		author string
		want   Book
		got    Book
	}{
		{
			name:   "1",
			author: "Пушкин",
			want:   Book{author: "Пушкин"},
		},
		{
			name:   "2",
			author: "Достоевский",
			want:   Book{author: "Достоевский"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetAuthor(tC.author)
			assert.Equal(t, tC.want.author, tC.got.author)
		})
	}
}

func Test_SetYear(t *testing.T) {
	testCases := []struct {
		name string
		year int
		want Book
		got  Book
	}{
		{
			name: "1",
			year: 2002,
			want: Book{year: 2002},
		},
		{
			name: "2",
			year: 1999,
			want: Book{year: 1999},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetYear(tC.year)
			assert.Equal(t, tC.want.year, tC.got.year)
		})
	}
}

func Test_SetSize(t *testing.T) {
	testCases := []struct {
		name string
		size int
		want Book
		got  Book
	}{
		{
			name: "1",
			size: 342,
			want: Book{size: 342},
		},
		{
			name: "2",
			size: 18,
			want: Book{size: 18},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetSize(tC.size)
			assert.Equal(t, tC.want.size, tC.got.size)
		})
	}
}

func Test_SetRate(t *testing.T) {
	testCases := []struct {
		name string
		rate float32
		want Book
		got  Book
	}{
		{
			name: "1",
			rate: 8.1,
			want: Book{rate: 8.1},
		},
		{
			name: "2",
			rate: 22.16,
			want: Book{rate: 22.16},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.SetRate(tC.rate)
			assert.Equal(t, tC.want.rate, tC.got.rate)
		})
	}
}

func Test_Compare(t *testing.T) {
	testCases := []struct {
		name  string
		books []Book
		comp  Comparator
		want  bool
		got   bool
	}{
		{
			name: "year",
			books: []Book{
				{144, "Горе от ума", "Грибоедов", 1825, 154, 4.6},
				{12, "Евгений Онегин", "Пушкин", 1830, 215, 4.5},
			},
			comp: Comparator{mode: YEAR},
			want: false,
		},
		{
			name: "size",
			books: []Book{
				{14, "Маленький принц", "Сент-Экзюпери", 1943, 220, 4.8},
				{123, "Муму", "Тургенев", 1854, 215, 4.9},
			},
			comp: Comparator{mode: SIZE},
			want: true,
		},
		{
			name: "rate",
			books: []Book{
				{12, "Евгений Онегин", "Пушкин", 1830, 215, 4.5},
				{14, "Маленький принц", "Сент-Экзюпери", 1943, 220, 4.8},
			},
			comp: Comparator{mode: RATE},
			want: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = tC.comp.Compare(&tC.books[0], &tC.books[1])
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
