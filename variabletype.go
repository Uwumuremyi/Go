package main

import (
	"fmt"
)

var y = 42

//Declare the variable with the identifier "z"
//is of type string
//and I assign the value "shaken, not stirred."

//STATIC PROGRAMMING LANGUAGE not a dynamic programming language

//A variable is declared to hold a value of a certain type

var z  string = "Shaken, not stirred"

var a string = `James said, "Shaken, not stirred"`

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Println("%T\n, a")

}