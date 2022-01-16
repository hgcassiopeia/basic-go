package main

import "fmt"

func main() {

	// strings variable declaration
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Luffy"
	nameThree = "Zoro"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi" //short hand version
	fmt.Println(nameFour)

	// int variable declaration
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 25 // can't assign (-) because this type start with 0
	numFour := 10

	fmt.Println(numOne, numTwo, numThree, numFour)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 888888888.7
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
