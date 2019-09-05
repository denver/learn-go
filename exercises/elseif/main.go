package main

import "fmt"

func main() {
	age := 101
	if age < 13 {
		fmt.Println("Wow, you're young")
	} else if age < 20 {
		fmt.Println("You're a teenager")
	} else if age < 30 {
		fmt.Println("You're in your 20s")
	} else if age < 40 {
		fmt.Println("you're in your 30s")
	} else if age < 50 {
		fmt.Println("You're getting there")
	} else {
		fmt.Println("You're over the hill")
	}
}
