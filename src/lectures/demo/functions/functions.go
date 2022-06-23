package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function!")
}
func main() {
	greet()
	dozen := double(6)
	fmt.Println("A dozen is", dozen)
	bakersDoze := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDoze)
	anotherBakersDoze := add(double(6), 1)
	fmt.Println("A baker's dozen is", anotherBakersDoze )
}
