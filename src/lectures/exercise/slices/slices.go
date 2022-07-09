//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAsseblyLine(asseblyLine []Part) {
	fmt.Println()
	fmt.Println("---Printing assembly line---")
	for i := 0; i < len(asseblyLine); i++ {
		part := asseblyLine[i]
		fmt.Println(part)
	}
}

func main() {
	asseblyLine := []Part{"Motor", "Oil tube", "Break disc"}
	printAsseblyLine(asseblyLine)
	asseblyLine = append(asseblyLine, "Break fluid", "Rotor")
	printAsseblyLine(asseblyLine)
	asseblyLine = asseblyLine[3:]
	printAsseblyLine(asseblyLine)
}
