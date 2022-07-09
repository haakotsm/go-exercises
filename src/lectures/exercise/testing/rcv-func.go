//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package rcv

import "fmt"

type Health int
type Energy int
type Name string

type Character struct {
	name      Name
	currentHP Health
	maxHP     Health
	currentEP Energy
	maxEP     Energy
}

func (char *Character) takeDamage(health Health) {
	if char.currentHP-health < 0 {
		char.currentHP = 0
	} else {
		char.currentHP -= health
	}
}
func (char *Character) gainHealing(health Health) {
	if char.currentHP+health > char.maxHP {
		char.currentHP = char.maxHP
	} else {
		char.currentHP += health
	}
}
func (char *Character) useEnergy(energy Energy) {
	if char.currentEP-energy < 0 {
		char.currentEP = 0
	} else {
		char.currentEP -= energy
	}
}
func (char *Character) gainEnergy(energy Energy) {
	if char.currentEP+energy > char.maxEP {
		char.currentEP = char.maxEP
	} else {
		char.currentEP += energy
	}
}

func (char *Character) displayStats() {
	fmt.Println()
	fmt.Println("----Stats----")
	fmt.Println("Name:", char.name)
	fmt.Println("Health:", char.currentHP, "/", char.maxHP)
	fmt.Println("Energy:", char.currentEP, "/", char.maxEP)
}
func main() {
	cloud := Character{"Cloud", 100, 100, 50, 50}
	cloud.displayStats()

	cloud.takeDamage(50)
	cloud.displayStats()

	cloud.useEnergy(10)
	cloud.displayStats()

	cloud.gainHealing(20)
	cloud.displayStats()

	cloud.gainEnergy(5)
	cloud.displayStats()
}
