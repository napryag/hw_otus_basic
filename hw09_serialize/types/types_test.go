package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JSON_Serialization(t *testing.T) {
	testCases := []struct {
		name  string
		buff  []byte
		examp Book
		got   Book
		want  Book
		err   error
	}{
		{
			name: "JSON Serialization 1",
			examp: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
			want: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
		},
		{
			name: "JSON Serialization 2",
			examp: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
			want: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.buff, tC.err = MarshalJSON(&tC.examp)
			UnmarshalJSON(tC.buff, &tC.got)
			assert.Equal(t, tC.want, tC.got, tC.err)
		})
	}
}

func Test_XML_Serialization(t *testing.T) {
	testCases := []struct {
		name  string
		buff  []byte
		examp Book
		got   Book
		want  Book
		err   error
	}{
		{
			name: "XML Serialization 1",
			examp: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
				Sample: []uint8{},
			},
			want: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
				Sample: []uint8{},
			},
		},
		{
			name: "XML Serialization 2",
			examp: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
				Sample: []uint8{},
			},
			want: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
				Sample: []uint8{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.buff, tC.err = MarshalXML(&tC.examp)
			UnmarshalXML(tC.buff, &tC.got)
			assert.Equal(t, tC.want, tC.got, tC.err)
		})
	}
}

func Test_YAML_Serialization(t *testing.T) {
	testCases := []struct {
		name  string
		buff  []byte
		examp Book
		got   Book
		want  Book
		err   error
	}{
		{
			name: "YAML Serialization 1",
			examp: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
				Sample: []uint8{},
			},
			want: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
				Sample: []uint8{},
			},
		},
		{
			name: "YAML Serialization 2",
			examp: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
				Sample: []uint8{},
			},
			want: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
				Sample: []uint8{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.buff, tC.err = MarshalYAML(&tC.examp)
			UnmarshalYAML(tC.buff, &tC.got)
			assert.Equal(t, tC.want, tC.got, tC.err)
		})
	}
}

func Test_GOB_Serialization(t *testing.T) {
	testCases := []struct {
		name  string
		buff  bytes.Buffer
		examp Book
		got   Book
		want  Book
		err   error
	}{
		{
			name: "GOB Serialization 1",
			examp: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
			want: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
		},
		{
			name: "GOB Serialization 2",
			examp: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
			want: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.buff, tC.err = EncodeGOB(&tC.examp)
			DecodeGOB(tC.buff, &tC.got)
			assert.Equal(t, tC.want, tC.got, tC.err)
		})
	}
}

func Test_BSON_Serialization(t *testing.T) {
	testCases := []struct {
		name  string
		buff  []byte
		examp Book
		got   Book
		want  Book
		err   error
	}{
		{
			name: "BSON Serialization 1",
			examp: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
			want: Book{
				ID:     18,
				Title:  "Дети родителей",
				Author: "Грибоводов",
				Year:   2024,
				Size:   2,
				Rate:   4.22,
			},
		},
		{
			name: "BSON Serialization 2",
			examp: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
			want: Book{
				ID:     22,
				Title:  "Домовой",
				Author: "Пушкин",
				Year:   1582,
				Size:   775,
				Rate:   4.2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.buff, tC.err = MarshalBSON(&tC.examp)
			UnmarshalBSON(tC.buff, &tC.got)
			assert.Equal(t, tC.want, tC.got, tC.err)
		})
	}
}
