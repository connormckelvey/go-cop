package main

import "fmt"

// VARS

var aSlice = []int{1, 2, 3, 4}

var aMap = map[string]int{
	"One":   1,
	"Two":   2,
	"Three": 3,
	"Four":  4,
}

var aChannel = make(chan int, 4)

func init() {
	for i := 0; i < 4; i++ {
		aChannel <- i + 1
	}
	close(aChannel)
}

// ENDVARS

func main() {
	for index, item := range aSlice {
		fmt.Printf("aSlice[%v] = %v \n", index, item)
	}

	for key, value := range aMap {
		fmt.Printf("aMap[%v] = %v \n", key, value)
	}

	for message := range aChannel {
		fmt.Printf("%v <-aChannel \n", message)
	}
}
