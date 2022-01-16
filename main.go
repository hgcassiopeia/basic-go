package main

import "fmt"

func main() {
	age := 17
	ageS := "17"
	name := "Luffy"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("New Line! \n")

	// Println
	fmt.Println("hello, with new line")
	fmt.Println("goodbye, with new line")
	fmt.Println("1: My age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("2: My age is %v and my name is %v \n", age, name)
	fmt.Printf("3: My age is %q and my name is %q \n", ageS, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("score is %f \n", 255.5)
	fmt.Printf("score is %0.1f \n", 255.5)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println(str)
}
