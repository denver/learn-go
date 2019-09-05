package main

import "fmt"

var gGroceries []string // slice of string

func addGrocery(a string) {
	fmt.Println("Capacity before:", cap(gGroceries))
	gGroceries = append(gGroceries, a)
	fmt.Println("Capacity after:", cap(gGroceries))
}

// func listGroceries() {
// 	fmt.Println("Grocery list is as follows:")
// 	for i := 0; i < len(gGroceries); i++ {
// 		fmt.Println(gGroceries[i])
// 	}
// }

func listGroceries() {
	for _, d := range gGroceries {
		fmt.Println(d)
	}
}

func main() {
	addGrocery("celery")
	addGrocery("coffee")
	addGrocery("juice")
	listGroceries()
	addGrocery("milk")
	addGrocery("butter")
	addGrocery("yogurt")
	listGroceries()
}
