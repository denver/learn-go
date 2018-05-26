package main

import "fmt"

var sum int = 0

func main() {
	for i := 1; i < 1000; i++ {
		if i%3 == 0 && i%5 == 0 {
			sum += i
		} else if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
