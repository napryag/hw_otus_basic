package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type val struct {
	hight, width int
	err          error
}

func Test_ScanValues(t *testing.T) {
	testCases := []struct {
		name         string
		hight, width int
		want         val
		got          val
	}{
		{
			name:  "целые положительные",
			hight: 8,
			width: 8,
			want:  val{8, 8, nil},
		},
		{
			name:  "целая отрицательная высота",
			hight: -8,
			width: 8,
			want:  val{-8, 8, errors.New("значение высоты должно быть положительным целым числом")},
		},
		{
			name:  "целая отрицательная ширина",
			hight: 8,
			width: -8,
			want:  val{8, -8, errors.New("значение ширины должно быть положительным целым числом")},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.hight, tC.got.width, tC.got.err = scanValues(tC.hight, tC.width)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_BoardWriter(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want string
		got  string
	}{
		{
			name: "четное число",
			num:  6,
			want: "#",
		},
		{
			name: "нечетное число",
			num:  5,
			want: " ",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = boardWriter(tC.num)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_BoardMaker(t *testing.T) {
	testCases := []struct {
		name string
		h, w int
		want [][]string
		got  [][]string
	}{
		{
			name: "equal",
			h:    8,
			w:    8,
			want: [][]string{
				{"#", " ", "#", " ", "#", " ", "#", " "},
				{" ", "#", " ", "#", " ", "#", " ", "#"},
				{"#", " ", "#", " ", "#", " ", "#", " "},
				{" ", "#", " ", "#", " ", "#", " ", "#"},
				{"#", " ", "#", " ", "#", " ", "#", " "},
				{" ", "#", " ", "#", " ", "#", " ", "#"},
				{"#", " ", "#", " ", "#", " ", "#", " "},
				{" ", "#", " ", "#", " ", "#", " ", "#"},
			},
		},
		{
			name: "not equal",
			h:    9,
			w:    4,
			want: [][]string{
				{"#", " ", "#", " "},
				{" ", "#", " ", "#"},
				{"#", " ", "#", " "},
				{" ", "#", " ", "#"},
				{"#", " ", "#", " "},
				{" ", "#", " ", "#"},
				{"#", " ", "#", " "},
				{" ", "#", " ", "#"},
				{"#", " ", "#", " "},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = boardMaker(tC.h, tC.w)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
