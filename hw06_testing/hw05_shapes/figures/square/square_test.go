package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSquare(t *testing.T) {
	testCases := []struct {
		name string
		side float64
		want Square
		got  Square
	}{
		{
			name: "1",
			side: 6,
			want: Square{Side: 6},
		},
		{
			name: "2",
			side: 13.2,
			want: Square{Side: 13.2},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got.NewSquare(tC.side)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}

func Test_CalculatePirimetr(t *testing.T) {
	testCases := []struct {
		name   string
		square Square
		want   float64
		got    float64
	}{
		{
			name:   "1",
			square: Square{Side: 6},
			want:   6 * 4,
		},
		{
			name:   "2",
			square: Square{Side: 13.2},
			want:   13.2 * 4,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = tC.square.CalculatePerimeter()
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
