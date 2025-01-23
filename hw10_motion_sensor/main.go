package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func sensor(ch chan<- int) {
	defer close(ch)
	max := big.NewInt(100)
	end := time.After(1 * time.Minute)

	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			continue
		}
		select {
		case <-end:
			return
		case ch <- int(n.Int64()):
		}
	}
}

func data(sch <-chan int, dch chan<- float64) {
	defer close(dch)
	var data float64
	var count int
	for val := range sch {
		data += float64(val)
		count++
		if count == 10 {
			dch <- data / 10
			data = 0
			count = 0
		}
	}
}

func main() {
	sensorCh := make(chan int)
	dataCh := make(chan float64)

	go sensor(sensorCh)
	go data(sensorCh, dataCh)

	for val := range dataCh {
		fmt.Println(val)
	}
}
