package main

import "fmt"

var wow int = 15

func testexit() {
	wow := 20
	fmt.Println(&wow, wow)
}

func main() {
	fmt.Println(&wow, wow)
	testexit()
	fmt.Println(&wow, wow)
}
