package main

import (
	"fmt"
	"sort"
)

type Widget struct {
	name   string
	weight int
}

// STARTBY

type byWeight []Widget

func (w byWeight) Len() int {
	return len(w)
}

func (w byWeight) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w byWeight) Less(i, j int) bool {
	return w[i].weight < w[j].weight
}

// ENDBY

func main() {
	widgets := []Widget{
		{name: "XL Widget", weight: 15},
		{name: "L Widget", weight: 10},
		{name: "S Widget", weight: 1},
		{name: "M Widget", weight: 4},
	}
	fmt.Println(widgets)

	sort.Sort(byWeight(widgets))
	fmt.Println(widgets)
}
