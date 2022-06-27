//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func checkFizz(num int) bool {
	return num%3 == 0
}
func checkBuzz(num int) bool {
	return num%5 == 0
}
func main() {
	for i := 1; i <= 50; i++ {
		if checkFizz(i) && checkBuzz(i) {
			fmt.Println("FizzBuzz")
		} else if checkFizz(i) {
			fmt.Println("Fizz")
		} else if checkBuzz(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
