package circle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCircle(t *testing.T) {
	testCases := []struct {
		name   string
		radius float64
		want   Circle
		got    Circle
	}{
		{
			name:   "1",
			radius: 6,
			want:   Circle{Radius: 6},
		},
		{
			name:   "2",
			radius: 13.2,
			want:   Circle{Radius: 13.2},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.NewCircle(tC.radius)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_CalculateArea(t *testing.T) {
	testCases := []struct {
		name   string
		circle Circle
		want   float64
		got    float64
	}{
		{
			name:   "1",
			circle: Circle{Radius: 6},
			want:   Pi * 6 * 6,
		},
		{
			name:   "2",
			circle: Circle{Radius: 13.2},
			want:   Pi * 13.2 * 13.2,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = tC.circle.CalculateArea()
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
