package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func subtract(in <-chan int, num int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n - num
		}
		close(out)
	}()
	return out
}

func someSlowFunc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			out <- n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	rand.Seed(time.Now().UTC().UnixNano())
	out := someSlowFunc(subtract(sq(gen(1, 2, 3, 4, 5, 6)), 8))

	for res := range out {
		fmt.Println(res)
	}
}
