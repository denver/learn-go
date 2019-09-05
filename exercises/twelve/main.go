package main

import "fmt"

func main() {

	fmt.Println("On the", 1, "day of Christmas my true love sent to me:")
	fmt.Println("a Partridge in a pear tree")
	fmt.Println("")

	for day := 2; day <= 12; day++ {
		fmt.Println("On the", day, "day of Christmas my true love gave to me:")
		switch day {
		case 12:
			fmt.Println("12 Drummers Drumming")
			fallthrough
		case 11:
			fmt.Println("11 Pipers Piping")
			fallthrough
		case 10:
			fmt.Println("10 Lors a Leaping")
			fallthrough
		case 9:
			fmt.Println("Nine Ladies Danciing")
			fallthrough
		case 8:
			fmt.Println("Eight Maids a Milking")
			fallthrough
		case 7:
			fmt.Println("7 Swans a Swimming")
			fallthrough
		case 6:
			fmt.Println("6 Geese of Laying")
			fallthrough
		case 5:
			fmt.Println("5 Golden Rings")
			fallthrough
		case 4:
			fmt.Println("4 calling birds")
			fallthrough
		case 3:
			fmt.Println("3 French Hens")
			fallthrough
		case 2:
			fmt.Println("2 Turtle Doves")
			fallthrough
		case 1:
			fmt.Println("and a Partride in a Pear Tree")
		}
		fmt.Println("")
	}
}
