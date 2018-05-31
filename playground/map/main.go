package main

import "fmt"

func main() {
	myMap := map[string]string{
		"k1": "key1",
		"k2": "key2",
		"k3": "key3",
	}

	for k, v := range myMap {
		fmt.Println(k, " = ", v)
	}
}
