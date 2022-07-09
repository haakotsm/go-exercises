//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package rcv

import "testing"

func TestMaxHealth(t *testing.T) {
	cloud := Character{"Cloud", 100, 100, 50, 50}
	cloud.gainHealing(10)
	if cloud.currentHP > cloud.maxHP {
		t.Errorf("Current health (%v) cannot exceed maximum health (%v)", cloud.currentHP, cloud.maxHP)
	}
}
func TestMaxEp(t *testing.T) {
	cloud := Character{"Cloud", 100, 100, 50, 50}
	cloud.gainEnergy(10)
	if cloud.currentEP > cloud.maxEP {
		t.Errorf("Current energy (%v) cannot exceed maximum energy (%v)", cloud.currentEP, cloud.maxEP)

	}
}

func TestMinHealth(t *testing.T) {
	cloud := Character{"Cloud", 100, 100, 50, 50}
	cloud.takeDamage(110)
	if cloud.currentHP < 0 {
		t.Errorf("Current health (%v) cannot go below 0", cloud.currentHP)
	}
}

func TestMinEP(t *testing.T) {
	cloud := Character{"Cloud", 100, 100, 50, 50}
	cloud.useEnergy(60)
	if cloud.currentEP < 0 {
		t.Errorf("Current energy (%v) cannot go below 0", cloud.currentEP)
	}
}
