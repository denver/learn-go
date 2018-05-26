package main

import "fmt"

func main() {
	fmt.Println(findGreatest(100, 1, 2, 3))
}

func findGreatest(nums ...int) int {
	var greatest int = 0
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}
