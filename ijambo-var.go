package main

import(
	"fmt"
)

var y = 43 //Scope of y will be package level

//Declare there is a variable with the identifier "z"
//and that the variable with the identifier "z" is of type int
//assigns the zero value of type int to "z"
var z int 

func main(){

	//Short declaration operator
	//Declare a variable and asign a value(of a certain type) at the same time.
	x := 42 
	fmt.Println( x )

	fmt.Println(y)

	foo()

	fmt.Print(z, "\n")
}

func foo(){
	fmt.Println("again:", y)
}