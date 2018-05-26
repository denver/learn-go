package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3}
	foo(aSlice...)
	foo()
}

func foo(n ...int) {
	for _, i := range n {
		fmt.Println(i)
	}
}
