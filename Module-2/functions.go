package main

//introduction to functions
//declaring functions
//function calling
//function arguments
//function demonstration

// go functions can return multiple values.

// You can call this function and assign the multiple return values to separate variables

func swap(a, b int) (int, int) {
	return b, a
}

// you can call this function and assign the multiple return values to separate variables.

x, y := swap(3, 5) // x=5, y=3

import "fmt"

func main() {
   

}

func add(a int, b int) int {
	return a + b
}

// call by value
// call by reference

// Variadic Functions:= In Go , you can create variadic functions that accepts a variable number of arguments. To declare a variadic function , you use an allipsis(...) before the last parameter

