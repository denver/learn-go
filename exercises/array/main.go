package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(b)
}
