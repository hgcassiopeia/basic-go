package main

import "fmt"

func main() {
	// string
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Luffy"
	nameThree = "Zoro"

	fmt.Println(nameOne, nameTwo, nameThree)
}
