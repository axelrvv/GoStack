package main

import (
	s "Stack/stack"
	"fmt"
)

type estudiante struct {
	name string
	nota int
	cell int
}

func main() {

	st := s.NewStack()

	a := estudiante{"Axel", 100, 8098437488}

	st.Push(a)

	x, err := st.Top()
	if err != nil {
		panic(err)
	}

	fmt.Println(x.(estudiante).name)
}
