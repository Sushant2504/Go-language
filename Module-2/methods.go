package main

import (
    "fmt"
)

// intro to methods
// declaring methods
// methods demonstration
// methods on built-in types
// embedded type and methods

// in go, a method is defined using the func keyword followed by the receiver type (the type on which the method operates) and a method name. The receiver tpe is specified within parenthesis before the method name. The syntax is demonstrated below.

type author struct{
	name string

    branch string

	particles int

	salary int
}

func (a author) show() {

	 fmt.println("Autor's name: ", a.name)

	 fmt.println("Branch name: ", a.branch)

}

type Myfloat float64

func (f MyFloat) Absolute() MyFloaot{
	 if f < 0 {
		return -f
	 }

	 return f
}

func main() {
	 
	add := func(a, b int) int{
		return a+b;
	}

	result := add(2,3)

	fmt.println("Result:", result)
   
}

func (int) Add(a, b int) int {
	 // method body

	 res := author(
		name: "Sushant",

		branch: "CSE",

		particles: 100,
		
        salary: 25000,
	 )

	 res.show()
}

//receiver: this is an identifier for the receiver type, and it can be any valid identifier,. it specifies on which type the method operates.
// ReceiverType: This is the data to which the method is attached. it can be a user-defined type or a built-in type.
//MethodName: These are the input values that the method expects, similiar to regular function parameters





