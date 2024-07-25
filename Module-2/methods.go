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

func main() {
	 

}

func (int) Add(a, b int) int {
	 // method body
}

//receiver: this is an identifier for the receiver type, and it can be any valid identifier,. it specifies on which type the method operates.
// ReceiverType: This is the data to which the method is attached. it can be a user-defined type or a built-in type.
//MethodName: These are the input values that the method expects, similiar to regular function parameters



