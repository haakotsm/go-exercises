package main

import "fmt"

type Space struct {
	occupied bool
}
type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func vacateSpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{make([]Space, 4)}
	fmt.Println("Initial:", lot)
	occupySpace(&lot, 1)
	fmt.Println("Second:", lot)
	lot.occupySpace(2)
	fmt.Println("Third:", lot)
	vacateSpace(&lot, 1)
	fmt.Println("Fourth:", lot)
	lot.vacateSpace(2)
	fmt.Println("Fifth:", lot)
}
