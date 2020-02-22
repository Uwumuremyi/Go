package main

import(
	"fmt"
)

var y string //zero value of y is an "empty string"
var z int //zero value of z (int) is 0

func main() {

	//Declare a variable to be of a certain type, and asign a value of
	//that type to the variable

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T", y)

	fmt.Println()

	y = "Bond, James"

	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T", y)

	fmt.Println()

	fmt.Println(z)
	fmt.Printf("%T", z)

	fmt.Println()


}