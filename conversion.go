package main

import "fmt"

// chapter five

type ShoppingList []string

func testConvertions() {
	list := ShoppingList{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	printSlice(slice)
	printShoppingList(list)
}

func printSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func printShoppingList(list ShoppingList) {
	fmt.Println("Shopping list: ", list)
}
