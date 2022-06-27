//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func check(shoppingList [4]Product) {
	length := 0
	sum := 0.00
	var lastItem Product
	for i := 0; i < len(shoppingList); i++ {
		product := shoppingList[i]
		if product.name != "" {
			lastItem = product
			length = i + 1
			sum += product.price
		} else {
			break
		}
	}
	fmt.Println("Last item on list", lastItem.name)
	fmt.Println(length)
	fmt.Println(sum)
}
func main() {
	shoppingList := [4]Product{
		{name: "Milk", price: 2.99},
		{name: "Bread", price: 1.99},
		{name: "Butter", price: 2.49},
	}
	check(shoppingList)

	shoppingList[3] = Product{name: "Ham", price: 4.00}

	check(shoppingList)
}
