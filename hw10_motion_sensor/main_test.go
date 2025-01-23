package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Sensor(t *testing.T) {
	senschan := make(chan int)
	go sensor(senschan)
	select {
	case num, ok := <-senschan:
		if !ok {
			t.Fatal()
		}
		assert.True(t, num >= 0 && num < 100)
	case <-time.After(1001 * time.Millisecond):
		t.Fatal()
	}
}

func Test_Data(t *testing.T) {
	testCases := []struct {
		name     string
		example  chan int
		example2 chan float64
		num      int
		got      float64
		want     float64
	}{
		{
			name:     "1",
			example:  make(chan int),
			example2: make(chan float64),
			num:      2,
			want:     11,
		},
		{
			name:     "2",
			example:  make(chan int),
			example2: make(chan float64),
			num:      12,
			want:     66,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			go func() {
				count := 0
				for {
					count++

					if count == 11 {
						return
					}
					tC.example <- tC.num * count
				}
			}()
			go data(tC.example, tC.example2)
			tC.got = <-tC.example2
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
