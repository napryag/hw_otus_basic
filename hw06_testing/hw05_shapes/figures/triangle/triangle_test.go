package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTriangle(t *testing.T) {
	testCases := []struct {
		name  string
		sides []float64
		want  Triangle
		got   Triangle
	}{
		{
			name:  "1",
			sides: []float64{10, 12},
			want:  Triangle{Hight: 10, Base: 12},
		},
		{
			name:  "2",
			sides: []float64{11.1, 22.2},
			want:  Triangle{Hight: 11.1, Base: 22.2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.NewTriangle(tC.sides[0], tC.sides[1])
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_CalculateArea(t *testing.T) {
	testCases := []struct {
		name     string
		triangle Triangle
		want     float64
		got      float64
	}{
		{
			name:     "1",
			triangle: Triangle{Hight: 6, Base: 8},
			want:     (6 * 8) / 2,
		},
		{
			name:     "2",
			triangle: Triangle{Hight: 18.8, Base: 20.2},
			want:     (18.8 * 20.2) / 2,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = tC.triangle.CalculateArea()
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
