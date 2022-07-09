package main

import "fmt"

type Preparer interface {
	PerpareDish()
}

func (c Chicken) PerpareDish() {
	fmt.Println("marinade")
	fmt.Println("cook")
}

func (s Salad) PerpareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PerpareDish()
	}
	fmt.Println()
}

type Chicken string
type Salad string

func main() {
	dishes := []Preparer{Chicken("Marinaded Chicken"), Salad("Salad")}
	prepareDishes(dishes)

}
