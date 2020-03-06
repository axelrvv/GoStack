package stack

import "strconv"

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

type strings struct {
	x string
}

type structs struct {
	x struct{}
}

var intsArr []Ints
var stringsArr []strings
var structsArr []struct{}

//Push : Add an int to the stack
func (i Ints) Push() string {
	number := strconv.Itoa(i.X)
	intsArr = append(intsArr, i)
	return "Pushed " + number
}

//Top : Show the las element inserted in the stack
func (i Ints) Top() string {
	length := len(intsArr)
	number := strconv.Itoa(intsArr[length-1].X)
	return number
}

//Pop : Delete the top element inserted in the stack
func (i Ints) Pop() string {
	length := len(intsArr)
	popped := strconv.Itoa(intsArr[length-1].X)
	intsArr = intsArr[:length-1]
	return popped
}

//IsEmpty : returns true is stack is empty else it returns false
func (i Ints) IsEmpty() bool {
	if len(intsArr) == 0 {
		return true
	}
	return false
}

//GetSize : retuns the size of the stack
func (i Ints) GetSize() string {
	size := len(intsArr)
	return strconv.Itoa(size)
}
