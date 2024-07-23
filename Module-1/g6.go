//file system in golang

package main

import "fmt"
import "os"


func main(){
	 file, err := os.open("sushant.txt")

	 if err != nil {
		 fmt.println("Error:", err)
		 return
	 }

	 defer file.close()

}

//concept of panic and recover

// the panic function is use to trigger a panic or runtime error explicity

func doSomething() {
	 
	if someCondition {
		 panic("An error occured")
	}

	//Rest of the code

}

//recover function is used to catch and handle panics

//panic and recover example:

func main() {

    defer func() {
         if r := recover(); r != nil {
             fmt.Println("Recovered in main:", r)
         }
    }()

    fmt.println("Before panic") 

	panic("Something went wrong!") 

	fmt.println("After panic")

}

func demonstratepanicAndRecover() {
	 fmt.println("\nPanic and Recover Example: ")

	 defer func() {
		 
		if r := recover(); r != nil {
			 fmt.Println("Recovered in demonstratepanicAndRecover:", r)
		}
	 }()

	 fmt.println("Before panic")

	 panic("Something went wrong!")
	 
	 fmt.Println("After panic")
}