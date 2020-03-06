package funcs

import "strconv"

type stack interface {
	//Add to stack
	push() string
	//Return last pushed
	top() string
	//Delete from stack the one on top
	pop() string
	//Says if stack is empty
	is_empty() bool
	//Says stack size
	get_size() int
}

type ints struct {
	x int
}

type strings struct {
	x string
}

type structs struct {
	x struct{}
}

var intsArr []ints
var stringsArr []strings
var structsArr []struct{}

func (i ints) push() string {
	number := strconv.Itoa(i.x)
	intsArr = append(intsArr, i)
	return "Pushed " + number
}

func (i ints) top() string {
	length := len(intsArr)

}

func (i ints) pop() string {
	return "TODO"
}
