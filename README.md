# GoStack

Stack data structure implemented in Go.

## Features

| **Feature** | **Status** | **Description** |
| ------------- | ------------- | ------------- |
|Push | ✅ | Adds an element to the stack |
| Pop | ✅ | Removes the top element of the stack |
| Top | ✅ |Returns the top element of the stack |
| IsEmpty | ✅ | Returns if stack is empty or not |
| GetSize | ✅ | Returns size of the stack |

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
}

/*
Output
--------
True
False
400
4
400
300
3
*/

```
*To make the list receptive for any datatype we deal with interfaces, it's important to know how to do interface assertion when you're storing a struct in the list*

### Example

```Go
package main

import (
	"fmt"

	st "github.com/axelrvv/GoStack"
)

type person struct {
	name   string
	age    int
	weight float32
}

func main() {

	stk := st.NewStack()

	fmt.Println(stk.IsEmpty())
	//Returns True

	p1 := person{"Axel Vasquez", 22, 180.88}
	p2 := person{"John Doe", 23, 150.50}
	p3 := person{"Doe John", 22, 140.78}

	stk.Push(p1)
	stk.Push(p2)
	stk.Push(p3)

	fmt.Println(stk.IsEmpty())
	//Returns false

	person1, err := stk.Top()

	if err != nil {
		panic(err)
	}

	name := person1.(person).name
	age := person1.(person).age
	weight := person1.(person).weight

	fmt.Printf("Name: %v; Age: %v; Weight: %v", name, age, weight)
}

/*
Output
---------------------------------------
true
false
Name: Doe John; Age: 22; Weight: 140.78
*/
```

# Download

go get github.com/axelrvv/GoStack

# A problem?

Don't hesitate to let me know, and make an issue.
