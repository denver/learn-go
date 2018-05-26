package main

import "fmt"

func main() {
	h, even := half(4)
	fmt.Println(h, even)
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
