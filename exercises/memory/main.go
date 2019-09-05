package main

import "fmt"

func main() {
	var a int32

	a = 27
	fmt.Println("Value of a", a)
	fmt.Println("Value of a", &a)
	fmt.Printf("%T", a)
}
