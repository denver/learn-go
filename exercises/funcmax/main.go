package main

import "fmt"

func max(i int, j int, k *int) {
	fmt.Println("value of k", k)
	if i > j {
		*k = i
	} else {
		*k = j
	}
}
func main() {
	var c int
	fmt.Println("address of c", &c)
	max(10, 15, &c)
	fmt.Println(c)
}
