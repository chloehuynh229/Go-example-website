package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Input a: ")

	// var then variable name then variable type
	var a float64

	// Taking input from user
	fmt.Scanln(&a)
	fmt.Println("Input b: ")
	var b float64
	fmt.Scanln(&b)

	// Print function is used to
	// display output in the same line
	// var x float64
	// fmt.Println("x =", -b/a)
	if a == 0 && b == 0 {
		fmt.Println("PT vo so nghiem .")
	} else if a == 0 && b != 0 {
		fmt.Println("Invalid")
		// // Addition of two string
		// fmt.Print(a + b)
	} else {
		fmt.Println("x =", -b/a)

	}
}
