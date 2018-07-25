package main

import (
	"fmt"
)

func loop() {
	sum := 0
	for {
		sum++
		fmt.Println(sum)
	}
}

func main() {
	loop()
}
