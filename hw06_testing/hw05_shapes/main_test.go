package main

import (
	"testing"

	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw05_shapes/figures/circle"
	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw05_shapes/figures/rectangle"
	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw05_shapes/figures/square"
	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw05_shapes/figures/triangle"
	"github.com/stretchr/testify/assert"
)

func Test_CalculateArea(t *testing.T) {
	testCases := []struct {
		name   string
		figure any
		err    error
		got    float64
		want   float64
	}{
		{
			name:   "triangle",
			figure: triangle.Triangle{Hight: 5, Base: 5},
			want:   float64(5*5) / 2,
		},
		{
			name:   "circle",
			figure: circle.Circle{Radius: 12},
			want:   circle.Pi * 12 * 12,
		},
		{
			name:   "rectangle",
			figure: rectangle.Rectangle{FirstSide: 32, SecondSide: 12},
			want:   32 * 12,
		},
		{
			name:   "square not implement interface",
			figure: square.Square{Side: 4},
			want:   0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got, tC.err = calculateArea(tC.figure)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
