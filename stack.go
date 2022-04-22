package main

import (
	"errors"
	"fmt"
)

// chapter three

type Stack struct {
	Values []interface{}
}

func testStack() {
	stack := Stack{}

	fmt.Println("Stack created with length ", stack.Length())
	fmt.Println("Is empty? ", stack.IsEmpty())

	stack.Push("Go")
	stack.Push(1997)
	stack.Push(3.14)
	stack.Push("The end")

	fmt.Println("Length after push 4 values: ", stack.Length())
	fmt.Println("Is empty? ", stack.IsEmpty())

	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		fmt.Println("Removing ", v)
		fmt.Println("Length ", stack.Length())
		fmt.Println("Is empty? ", stack.IsEmpty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func (stack Stack) Length() int {
	return len(stack.Values)
}

func (stack Stack) IsEmpty() bool {
	return stack.Length() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.Values = append(stack.Values, value)
}

func (stack *Stack) Pop() (interface{}, error) {

	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	value := stack.Values[stack.Length()-1]
	stack.Values = stack.Values[:stack.Length()-1]

	return value, nil
}
