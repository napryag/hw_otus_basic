package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StrEditor(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		got  []string
		want []string
	}{
		{
			name: "1",
			str:  "Весь полет она провела, ходя взад и вперед между четырьмя туалетами на борту.",
			want: []string{"весь", "полет", "она", "провела", "ходя", "взад", "и", "вперед",
				"между", "четырьмя", "туалетами", "на", "борту"},
		},
		{
			name: "2",
			str:  "Сначала все подумали, что она заболела, но позже экипаж понял, что она не входит в число пассажиров.",
			want: []string{"сначала", "все", "подумали", "что", "она", "заболела", "но",
				"позже", "экипаж", "понял", "что", "она", "не", "входит", "в", "число", "пассажиров"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = strEditor(tC.str)
			assert.Equal(t, tC.got, tC.want)
		})
	}
}

func Test_CountWords(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		got  map[string]int
		want map[string]int
	}{
		{
			name: "1",
			str:  "Весь полет она провела, ходя взад и вперед между четырьмя туалетами на борту.",
			want: map[string]int{
				"весь": 1, "полет": 1, "она": 1, "провела": 1, "ходя": 1, "взад": 1, "и": 1,
				"вперед": 1, "между": 1, "четырьмя": 1, "туалетами": 1, "на": 1, "борту": 1,
			},
		},
		{
			name: "2",
			str:  "Сначала все подумали, что она заболела, но позже экипаж понял, что она не входит в число пассажиров.",
			want: map[string]int{
				"сначала": 1, "все": 1, "подумали": 1, "что": 2, "она": 2, "заболела": 1, "но": 1,
				"позже": 1, "экипаж": 1, "понял": 1, "не": 1, "входит": 1, "в": 1, "число": 1, "пассажиров": 1,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got = countWords(tC.str)
			assert.Equal(t, tC.got, tC.want)
		})
	}
}
