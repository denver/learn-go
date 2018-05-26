package main

import "fmt"

func main() {
	fib1 := 0
	fib2 := 1
	fib3 := 0
	sum := 0
	for fib1+fib2 < 4000000 {
		fib3 = fib1 + fib2
		fmt.Println(fib3)
		if fib3%2 == 0 {
			sum += fib3
		}
		fib1 = fib2
		fib2 = fib3
	}
	fmt.Println("Sum of evens", sum)
}
