package main

import "fmt"

func main() {

   // pointer data types 

   var num int = 42;

   var ptr *int

   ptr = &num

   fmt.println("Value of num: %d\n", num)

   fmt.printf("Value pointed to by ptr: %d\n", *ptr)



}