package cmd

import (
	"fmt"

	"../pkg/stack"
)

func StackMain() {
	stack := &stack.StringListStack{}
	stack.Push("kim")
	fmt.Printf("%+v", stack.Head)
	stack.Push("lip")
	fmt.Printf("%+v", stack.Head)
	d := stack.Pop()
	fmt.Printf("%+v", d)
	fmt.Println(stack.IsEmpty())
	d = stack.Pop()
	fmt.Println(stack.IsEmpty())
}
