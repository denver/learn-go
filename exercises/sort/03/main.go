package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 8, 4, 1, 4, 10, 1111, 34324, 23, 23}
	sort.Ints(n)
	fmt.Println(n)
}
