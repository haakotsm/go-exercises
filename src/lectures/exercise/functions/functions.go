//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Printf("Greetings %v!\n", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func anyMsg() string {
	return "Any Message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func sum3Numbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func randomNumber() int {
	return rand.Int()
}

//* Write a function that returns any two numbers
func twoRandomNumbers() (int, int) {
	return randomNumber(), randomNumber()
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once
func main() {
	greet(anyMsg())
	val1, val2 := twoRandomNumbers()
	val3 := randomNumber()
	fmt.Println(val1, "+", val2, "+", val3, "=", sum3Numbers(val1, val2, val3))
}
