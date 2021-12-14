package main

import "fmt"

func main() {
	stack := Stack{Size: 2}

	stack.Push("Tiago")
	stack.Push("Dani")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
