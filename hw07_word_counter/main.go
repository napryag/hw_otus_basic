package main

import (
	"fmt"
	"strings"
)

func strEditor(s string) []string {
	s = strings.Join(strings.Fields(s), " ")
	rep := []string{",", ".", "!", "?", ":", ";", "'", "@", "#", "$", "%", "№"}
	for key := range rep {
		s = strings.ReplaceAll(s, rep[key], "")
	}
	s = strings.ToLower(s)
	sl := strings.Split(s, " ")
	return sl
}

func countWords(text string) map[string]int {
	str := strEditor(text)
	mp := map[string]int{}

	for _, s := range str {
		mp[s]++
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}
	return mp
}

func main() {
	bb := countWords("      Теперь  поговорим  немного!   о  том,!?  как  пакет  strings  может увеличить производительность может как о Теперь. Теперь.")
	fmt.Println(bb["как"])
	fmt.Println(bb["теперь"])
}
