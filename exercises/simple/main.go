package main

import "fmt"

var gGroceries []string // slice of string

func addGrocery(a ...string) {
	for _, d := range a {
		gGroceries = append(gGroceries, d)
	}
}

func listGroceries() {
	for _, d := range gGroceries {
		fmt.Println(d)
	}
}

func main() {
	addGrocery("celery", "broccoli", "onions")
	addGrocery("coffee")
	addGrocery("juice")
	addGrocery("milk")
	addGrocery("butter")
	addGrocery("yogurt")
	listGroceries()
}
