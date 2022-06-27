//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(numberOfDice, dieSize, numberOfRolls int) {
	if numberOfDice < 1 || dieSize < 1 || numberOfRolls < 1 {
		fmt.Println("All values must be larger than 1.")
		return
	}
	counter := 0
	start := time.Now()
	for i := 0; i < numberOfRolls; i++ {
		// for {

		sum := 0
		for j := 0; j < numberOfDice; j++ {
			rand.Seed(time.Now().UnixNano())
			sum += (rand.Intn(dieSize) + 1)
		}
		fmt.Println(sum)
		isSnakeEye := numberOfDice == sum
		if isSnakeEye {
			end := time.Now()
			fmt.Println(sum, "is snakeyes. Found on iteration", counter, "after", end.Sub(start).Seconds(), "seconds")
		}
		isLucky := sum == 7
		if isLucky {
			fmt.Println(sum, "is lucky.")
		}
		isEven := sum%2 == 0
		if isEven {
			fmt.Println(sum, "is even")
		} else {
			fmt.Println(sum, "is odd")
		}
		counter++

	}
}
func main() {
	rollDice(7, 6, 10)
}
