package main

import "fmt"

func main() {
	greet(a)
	greet(b)
	greet("denver")
}

func greet(name string) {
	fmt.Println(name)
}

var a = "bob"
var b = "sally"

greet("tim")
