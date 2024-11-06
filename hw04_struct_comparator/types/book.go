package types

type enum int

const (
	YEAR enum = iota + 1
	SIZE
	RATE
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBook(id int, titile, author string, year, size int, rate float32) *Book {
	return &Book{
		id:     id,
		title:  titile,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

type Comparator struct {
	mode enum
}

func NewComparator(mode enum) *Comparator {
	return &Comparator{
		mode: mode,
	}
}

func (c Comparator) Compare(b1, b2 Book) bool {
	switch c.mode {
	case YEAR:
		return b1.year > b2.year
	case SIZE:
		return b1.size > b2.size
	case RATE:
		return b1.rate > b2.rate
	default:
		return false
	}
}
