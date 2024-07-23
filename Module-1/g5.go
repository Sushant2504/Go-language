package main

import "fmt"
import "os" // it is used to tapping into operating system calls



func main() {
	 
	for i:= 0; i<5; i++ {
		 println("*")
	}

    // infinite loop , without any condition

	for{
		fmt.println("Welcome to the infinite loop")
	}
 

	//break and continue statements

    // goto statement: the goto statement is a control statement that allows you to unconditionally jump to a labeled section of your code

	// Defer statement: It is used in panic and recover statements in go

    
}


//defer statement

func someFunction() {

	 defer cleanup()
    
	 //function body
	  
}