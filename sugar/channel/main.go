package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	// Syntaxt Reminds me of an IIFE in JS, similiar scoping
	// (function(input) {
	//	console.log(input)
	// })(input)

	go func() {
		for i := 0; i < 4; i++ {
			messages <- fmt.Sprintf("Hello from unit %v", i)
		}
		close(messages)
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
