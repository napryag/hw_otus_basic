package main

import "fmt"

func binarySearch(storage []int, s int) int {
	minIndex := 0
	maxIndex := len(storage) - 1
	for minIndex <= maxIndex {
		midIndex := (minIndex + maxIndex) / 2

		if s == storage[midIndex] {
			return s
		} else if s < storage[midIndex] {
			maxIndex = midIndex - 1
		} else {
			minIndex = midIndex + 1
		}
	}
	return -1
}

func main() {
	st := []int{11, 22, 33, 42, 54, 65, 78, 92, 99}
	fmt.Println(binarySearch(st, 65))
}
