package main

import "fmt"

const gCap int = 5          // total number of array elements
var gLen int                // integer number of items in our array currently
var gGroceries [gCap]string // array of strings

func add(a string) {
	if gLen < gCap {
		gGroceries[gLen] = a
		gLen++
	} else {
		fmt.Println("Error!!! We've reach the upper bound of the array!!!")
	}
}

func listGroceries() {
	fmt.Println("Grocery List is as follows:")
	for i := 0; i < gLen; i++ {
		fmt.Println(gGroceries[i])
	}
}

func main() {
	listGroceries()
	fmt.Println(gLen)
	add("coffee")
	fmt.Println(gLen)
	add("milk")
	add("fries")
	add("pizza")
	add("beer")
	add("broccoli")
	add("beets")
	fmt.Println(gLen)
	listGroceries()
}
