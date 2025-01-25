package main

import (
	"fmt"
	"sync"
)

func worker(num int) int {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
	)
	for i := 1; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Printf("Воркер %d: значение счетчика увеличено на 1\n", i)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return counter
}

func main() {
	counter := worker(754)
	fmt.Printf("Итоговое значение счетчика равно %d", counter)
}
