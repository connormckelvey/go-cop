package main

import (
	"fmt"
)

// START
type Person struct {
	name string
}

type Employee struct {
	Person
	number int
}

// END

func main() {
	e := Employee{
		Person{
			name: "Connor",
		},
		102021,
	}

	fmt.Println(e.name)
}
