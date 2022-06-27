//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Point struct {
	x, y int
}

type Rectangle struct {
	width, length int
}

func getArea(shape Rectangle) int {
	return shape.length * shape.width
}

func getPerimeter(shape Rectangle) int {
	return (shape.length * 2) + (shape.width * 2)
}

func getAreaAndPerimeter(shape Rectangle) (int, int) {
	return getArea(shape), getPerimeter(shape)
}

func main() {
	smallBox := Rectangle{10, 10}
	bigBox := Rectangle{length: 100, width: 50}

	smallArea, smallPerimeter := getAreaAndPerimeter(smallBox)
	bigArea, bigPerimeter := getAreaAndPerimeter(bigBox)
	fmt.Println("Small box area", smallArea, "and perimiter", smallPerimeter)
	fmt.Println("Big box area", bigArea, "and perimiter", bigPerimeter)
	smallBox.length *= 2
	smallBox.width *= 2
	newSmallArea, newSmallPerimeter := getAreaAndPerimeter(smallBox)
	fmt.Println("Small box area", newSmallArea, "and perimiter", newSmallPerimeter)
}
