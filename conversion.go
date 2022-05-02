package main

import "fmt"

// chapter five

type ShoppingList []string
type GenericList []interface{}

func testConvertions() {
	list := ShoppingList{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	printSlice(slice)
	printShoppingList(list)
}

func testGenerics() {
	list := GenericList{1, "Coffee", 42, true, 23, "Ball", 3.14, false}

	fmt.Println(list.removeFisrt())
	fmt.Println(list.removeLast())
	fmt.Println(list.removeByIndex(3))
}

func printSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func printShoppingList(list ShoppingList) {
	fmt.Println("Shopping list: ", list)
}

func (list *GenericList) removeByIndex(index int) interface{} {
	l := *list
	removed := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return removed
}

func (list *GenericList) removeFisrt() interface{} {
	return list.removeByIndex(0)
}

func (list *GenericList) removeLast() interface{} {
	return list.removeByIndex(len(*list) - 1)
}
