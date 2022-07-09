//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type ModelType int
type ModelName string

type Lifter interface {
	Lift()
}

const (
	Motorcycle ModelType = iota
	Car
	Truck
)

func (m ModelType) String() string {
	switch m {
	case Motorcycle:
		return "Motorcycle"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	}
	return ""
}

type Vehicle struct {
	modelType ModelType
	modelName ModelName
}

type Shop string

func (v Vehicle) Lift() {
	fmt.Println("Checking vehicle type")
	switch v.modelType {
	case Motorcycle:
		fmt.Printf("%v, type %v goes to the small lift\n", v.modelName, v.modelType)

	case Car:
		fmt.Printf("%v, type %v goes to the standard lift\n", v.modelName, v.modelType)

	case Truck:
		fmt.Printf("%v, type %v goes to the large lift\n", v.modelName, v.modelType)
	}
	fmt.Println()
}

func (shop Shop) repairVehicle(lifters []Lifter) {
	for i := 0; i < len(lifters); i++ {
		lifter := lifters[i]
		lifter.Lift()
	}
}

func main() {
	shop := Shop("Riis Bilglass")
	car := Vehicle{Car, "Ford Mondeo"}
	truck := Vehicle{Truck, "Ford Ranger"}
	motorcycle := Vehicle{Motorcycle, "Kawazaki Ninja"}
	lifters := []Lifter{car, truck, motorcycle}
	shop.repairVehicle(lifters)
}
