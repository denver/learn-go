package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"zeno", "john", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)
}
