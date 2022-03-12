package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	concurrency := 8
	in := make(chan int)
	done := make(chan []byte)

	go func() {
		i := 0
		for {
			in <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go ProcessWorker(in, x)
	}
	<-done
}

func ProcessWorker(in chan int, worker int) {
	for x := range in {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker", worker, ": processing", x)
	}
}