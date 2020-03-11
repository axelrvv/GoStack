# GoStack
All functions of a stack data structure

## Features
- Push --> Add to the stack
- Pop --> Delete Top element of the stack
- Top --> Returns top element of the stack
- IsEmpty --> Returns if stack is empty
- GetSize --> Returns size of the stack

Now you can use them in go

# Usage
```Go
package main

import (
	"fmt"

	st "github.com/axelrvv/GoStack"
)

func main() {

	stk := st.NewStack()

	fmt.Println(stk.IsEmpty())
	//Returns True

	stk.Push(100)
	stk.Push(200)
	stk.Push(300)
	stk.Push(400)

	fmt.Println(stk.IsEmpty())
	//Returns false

	Top, err := stk.Top()
	if err != nil {
		panic(err)
	}
	fmt.Println(Top)
	//Returns 400
	fmt.Println(stk.GetSize())
	//Returns 4

	Popped, _ := stk.Pop()
	fmt.Println(Popped)
	//Returns 400

	Top, _ = stk.Top()
	fmt.Println(Top)
	//Returns 300

	fmt.Println(stk.GetSize())
	//Returns 3

	//Output
	// True
	// False
	// 400
	// 4
	// 400
	// 300
	// 3
```
# Download
go get github.com/axelrvv/GoStack
