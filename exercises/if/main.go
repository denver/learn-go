package main

import "fmt"

func main() {
	age := 41
	if age < 13 {
		fmt.Println("Wow, you're young")
	}

	if age > 12 && age < 20 {
		fmt.Println("You're a teenager")
	}

	if age > 20 && age < 30 {
		fmt.Println("You're in your 20s")
	}

	if age > 29 && age < 40 {
		fmt.Println("you're in your 30s")
	}

	if age > 40 && age < 50 {
		fmt.Println("You're getting there")
	}

	if age > 49 {
		fmt.Println("You're over the hill")
	}

}
