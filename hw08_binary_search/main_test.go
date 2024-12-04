package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearch(t *testing.T) {
	testCases := []struct {
		name    string
		storage []int
		find    int
		got     int
		want    int
	}{
		{
			name:    "midVal",
			storage: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			find:    5,
			want:    5,
		},
		{
			name:    "existingVal",
			storage: []int{12, 22, 54, 67, 90, 112, 128, 134, 167, 184, 211, 222, 231},
			find:    222,
			want:    222,
		},
		{
			name:    "notExistingVal",
			storage: []int{4, 7, 10, 33, 35, 38, 39, 44},
			find:    3,
			want:    -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = binarySearch(tC.storage, tC.find)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
