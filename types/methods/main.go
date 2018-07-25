package main

import (
	"fmt"
)

// STARTTYPES
type Inches int

// Define Methods on Any Named Type in your Package
func (i Inches) String() string {
	return fmt.Sprintf("%v'%v\"", int(i)/12, int(i)%12)
}

type Person struct {
	firstName string
	lastName  string
	height    Inches
}

func (p *Person) Name() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}

func (p *Person) Height() string {
	return p.height.String()
}

// ENDTYPES

func main() {
	p := &Person{
		firstName: "Connor",
		lastName:  "McKelvey",
		height:    71,
	}

	fmt.Println(p.Name())
	fmt.Println(p.Height())
}
