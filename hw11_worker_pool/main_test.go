package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
		got  int
		want int
	}{
		{
			desc: "1",
			num:  847,
			want: 846,
		},
		{
			desc: "2",
			num:  9083,
			want: 9082,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.got = worker(tC.num)
			assert.Equal(t, tC.got, tC.want)
		})
	}
}
