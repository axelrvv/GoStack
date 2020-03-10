package stack

import (
	"errors"
	"strconv"
)

type funcs interface {
	//Add to stack
	Push()
	//Return last pushed
	Top()
	//Delete from stack the one on top
	Pop()
	//Says if stack is empty
	Isempty()
	//Says stack size
	Getsize()
}

//Ints : int data type
type Ints struct {
	X int
}

//Strings : string data type
type Strings struct {
	X string
}

var intsArr []Ints
var stringsArr []Strings

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

//Push : Add an int to the stack
func (i Ints) Push() string {
	number := strconv.Itoa(i.X)
	intsArr = append(intsArr, i)
	return "Pushed " + number
}

//Push : Add a string to the stack
func (i Strings) Push() string {
	stringsArr = append(stringsArr, i)
	return stringsArr[len(stringsArr)-1].X
}

//IsEmpty : returns true is stack is empty else it returns false
func (i *Stack) IsEmpty() bool {
	if len(i.data) == 0 {
		return true
	}

	return false
}

//IsEmpty : returns true is stack is empty else it returns false
func (i Ints) IsEmpty() bool {
	if len(intsArr) == 0 {
		return true
	}
	return false
}

//IsEmpty : returns true is stack is empty else it returns false
func (i Strings) IsEmpty() bool {
	if len(stringsArr) == 0 {
		return true
	}
	return false
}

//Top : Show the las element inserted in the stack
func (i *Stack) Top() (Elem, error) {
	if i.IsEmpty() {
		return "", errors.New("Stack is empty")
	}
	return i.data[len(i.data)-1], nil
}

//Top : Show the las element inserted in the stack
func (i Ints) Top() (string, error) {
	if Ints.IsEmpty(i) {
		return "", errors.New("Stack is empty")
	}
	length := len(intsArr)
	number := strconv.Itoa(intsArr[length-1].X)
	return number, nil
}

//Top : Show the las element inserted in the stack
func (i Strings) Top() (string, error) {
	if Strings.IsEmpty(i) {
		return "", errors.New("Stack is empty")
	}
	top := stringsArr[len(stringsArr)-1].X
	return top, nil
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

//Pop : Delete the top element inserted in the stack
func (i Ints) Pop() (string, error) {
	if Ints.IsEmpty(i) {
		return "", errors.New("Stack is empty")
	}
	length := len(intsArr)

	popped := strconv.Itoa(intsArr[length-1].X)
	intsArr = intsArr[:length-1]
	return popped, nil
}

//Pop : Delete the top element inserted in the stack
func (i Strings) Pop() (string, error) {
	if Strings.IsEmpty(i) {
		return "", errors.New("Stack is empty")
	}
	popped := stringsArr[len(stringsArr)-1].X
	stringsArr = stringsArr[:len(stringsArr)-1]
	return popped, nil
}

//GetSize : returns the size of the stack
func (i *Stack) GetSize() int {
	size := len(i.data)
	return size
}

//GetSize : returns the size of the stack
func (i Ints) GetSize() int {
	size := len(intsArr)
	return size
}

//GetSize : returns the size of the stack
func (i Strings) GetSize() int {
	size := len(stringsArr)
	return size
}
