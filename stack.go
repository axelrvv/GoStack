package stack

import (
	"errors"
)

// type funcs interface {
// 	//Add to stack
// 	Push()
// 	//Return last pushed
// 	Top()
// 	//Delete from stack the one on top
// 	Pop()
// 	//Says if stack is empty
// 	Isempty()
// 	//Says stack size
// 	Getsize()
// }

//Elem : Is going to be where we store our data
type Elem interface{}

//Stack : Defining the data type
type Stack struct {
	data []Elem
}

//NewStack : To start a stack
func NewStack() Stack {
	st := Stack{}
	return st
}

//Push : Add a data to the stack
func (i *Stack) Push(e Elem) {
	i.data = append(i.data, e)
}

//IsEmpty : returns true is stack is empty else it returns false
func (i *Stack) IsEmpty() bool {
	if len(i.data) == 0 {
		return true
	}

	return false
}//This is to test circle ci
//Top : Show the las element inserted in the stack
func (i *Stack) Top() (Elem, error) {
	if i.IsEmpty() {
		return "", errors.New("Stack is empty")
	}
	return i.data[len(i.data)-1], nil
}

//Pop : Delete the top element inserted in the stack
func (i *Stack) Pop() (Elem, error) {
	if i.IsEmpty() {
		return "", errors.New("Stack is empty")
	}
	popped := i.data[len(i.data)-1]
	i.data = i.data[:len(i.data)-1]
	return popped, nil
}

//GetSize : returns the size of the stack
func (i *Stack) GetSize() int {
	size := len(i.data)
	return size
}
