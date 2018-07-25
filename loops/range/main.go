package main

import (
	"fmt"
)

func loop() {
	slice := []string{"these", "are", "strings", "in", "a", "slice"}
	for i, item := range slice {
		fmt.Printf("slice[%v] is: %v \n", i, item)
	}
}

func main() {
	loop()
}
