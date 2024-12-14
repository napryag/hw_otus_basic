package rectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRectangle(t *testing.T) {
	testCases := []struct {
		name  string
		sides []float64
		want  Rectangle
		got   Rectangle
	}{
		{
			name:  "1",
			sides: []float64{10, 12},
			want:  Rectangle{FirstSide: 10, SecondSide: 12},
		},
		{
			name:  "2",
			sides: []float64{111, 222},
			want:  Rectangle{FirstSide: 111, SecondSide: 222},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.NewRectangle(tC.sides[0], tC.sides[1])
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_CalculateArea(t *testing.T) {
	testCases := []struct {
		name      string
		rectangle Rectangle
		want      float64
		got       float64
	}{
		{
			name:      "1",
			rectangle: Rectangle{FirstSide: 6, SecondSide: 8},
			want:      6 * 8,
		},
		{
			name:      "2",
			rectangle: Rectangle{FirstSide: 18.8, SecondSide: 20.2},
			want:      18.8 * 20.2,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = tC.rectangle.CalculateArea()
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
