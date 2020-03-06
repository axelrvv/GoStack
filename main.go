package main

import (
	s "Stack/stack"
	"fmt"
	"strconv"
)

func main() {

	t := s.Ints{X: 6}

	s.Ints.Push(t)

	t.X = 4
	s.Ints.Push(t)

	t.X = 1
	s.Ints.Push(t)

	t.X = 9
	s.Ints.Push(t)

	u := s.Ints{X: 120}

	s.Ints.Pop(t)
	s.Ints.Push(u)
	fmt.Println("Top: " + s.Ints.Top(u))
	fmt.Println("Is it empty? " + strconv.FormatBool(s.Ints.IsEmpty(t)))
	fmt.Println("Size: " + s.Ints.GetSize(u))

}
