// introduction to array
// declaring arrays
// operations on arrays

// arrays in functions
// arrays example
// important points about arrays

// arrays have fixed size
// arrays are homogeneous
// arrays are value type in go
// go uses zero=based indexing in arrays

// arrays in the functions

package main

import "fmt"

func main() {

	// declaring an array of 5 integers
	var numbers [5]int

	// assigning values to the array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// accessing values from the array
	fmt.Println(numbers[0]) // Output: 10
	fmt.Println(numbers[1]) // Output: 20
	fmt.Println(numbers[2]) // Output: 30
	fmt.Println(numbers[3]) // Output: 40
	fmt.Println(numbers[4]) // Output: 50

	var my_slice = []string{"hello", "Dear", "Learners"}

	fmt.println("My Slice 1:", my_slice)

	my_slice_2 := []int{2, 54, 76, 57, 32, 14, 40}

	fmt.println("My slice 2:", my_slice_2)

	// slice using make() demonstration

	// the example represents intialising and creating slices using arrays

	var my_slice_1 = make([]int, 4, 7)

	fmt.printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n", my_slice2, len(my_slice_2), cap(my_slice_2))

	// accessing values out of the range of the array
	//fmt.Println(numbers[
}

// slices in the golang
// components of slices
// declaring slices
// creating slices
// intialising slices

// componets of slice:
// 1) pointers
// 2) length
// 3) capacity

// slice is used with : operator
// slice literal demonstration
// slices using arrays demonstration
// slices using make() demonstration
