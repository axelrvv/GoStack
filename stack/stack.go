package stack

import (
	"errors"
	"strconv"
)

type stack interface {
	//Add to stack
	Push() string
	//Return last pushed
	Top() string
	//Delete from stack the one on top
	Pop() string
	//Says if stack is empty
	Isempty() bool
	//Says stack size
	Getsize() int
}

//Ints : int data type
type Ints struct {
	X int
}

//Strings : string data type
type Strings struct {
	X string
}

//Data : Any data type can be stored here
type Data struct {
	X interface{}
}

var intsArr []Ints
var stringsArr []Strings
var dataArr []Data

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

//GetSize : retuns the size of the stack
func (i Ints) GetSize() int {
	size := len(intsArr)
	return size
}

//GetSize : returns the size of the stack
func (i Strings) GetSize() int {
	size := len(stringsArr)
	return size
}
