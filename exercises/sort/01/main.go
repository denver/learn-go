package main

import (
	"fmt"
	"sort"
)

func main() {

	type people []string
	studyGroup := people{"Zemo", "John", "Al", "Jenny"}

	if sort.StringsAreSorted(studyGroup) {
		fmt.Println("Sorted")
	} else {
		sort.Strings(studyGroup)
		fmt.Println(studyGroup)
		fmt.Println("not yet buddy")
	}
}
