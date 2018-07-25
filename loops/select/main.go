package main

import (
	"fmt"
)

func loop() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Message Number %v", i)
		}
		close(ch)
	}()

	for {
		fmt.Println("Select to Follow...")

		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(msg)
		}
	}
}

func main() {
	loop()
}
