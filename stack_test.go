package stack

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewStack(t *testing.T) {
	st := NewStack()
	test := len(st.data)
	if test != 0 {
		t.Errorf("st.data len should be %v", st.data)
	}
}

func TestPush(t *testing.T) {
	st := NewStack()
	tests := []struct {
		x      interface{}
		wanted int
	}{
		{10, 1},
		{"Test", 2},
		{true, 3},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%v", test.x, test.wanted)
		t.Run(testname, func(t *testing.T) {
			st.Push(test.x)
			if len(st.data) != test.wanted {
				t.Errorf("st.data len should be %v", test.wanted)
			}
		})
	}
}

func TestPopWhenStactEmpty(t *testing.T) {
	st := NewStack()
	el, _ := st.Pop()
	if el != "" {
		t.Errorf("Exepected an empty string")
	}
}

func TestPop(t *testing.T) {
	st := NewStack()

	tests := []struct {
		x         interface{}
		wantedLen int
	}{
		{10, 2},
		{"Test", 1},
		{true, 0},
	}

	for _, test := range tests {
		st.Push(test.x)
	}
	for index, test := range tests {
		testname := "Pop " + strconv.Itoa(index)
		t.Run(testname, func(t *testing.T) {
			st.Pop()
			if len(st.data) != test.wantedLen {
				t.Errorf("Expected %v", test.wantedLen)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	st := NewStack()
	if st.IsEmpty() != true {
		t.Errorf("Expectected %v", true)
	}
}

func TestTopWhenStackEmpty(t *testing.T) {
	st := NewStack()
	test, _ := st.Top()

	if test != "" {
		t.Error("Expected an empty string")
	}
}

func TestTop(t *testing.T) {
	st := NewStack()

	tests := []struct {
		x      interface{}
		wanted interface{}
	}{
		{10, 10},
		{"Test", "Test"},
		{true, true},
	}

	for _, test := range tests {
		st.Push(test.x)
		top, _ := st.Top()
		if top != test.wanted {
			t.Errorf("Expected %v", test.x)
		}
	}
}

func TestGetSizeWhenStackEmpty(t *testing.T) {
	st := NewStack()
	size := st.GetSize()
	if size != 0 {
		t.Error("Stack is supposed to be empty")
	}
}

func TestGetSize(t *testing.T) {
	st := NewStack()

	tests := []struct {
		x    interface{}
		size int
	}{
		{10, 1},
		{"Test", 2},
		{true, 3},
	}

	for _, test := range tests {
		st.Push(test.x)
		if st.GetSize() != test.size {
			t.Errorf("Expected %v", test.size)
		}
	}
}

/*
Output

=== RUN   TestNewStack
--- PASS: TestNewStack (0.00s)
=== RUN   TestPush
=== RUN   TestPush/10,1
=== RUN   TestPush/%!d(string=Test),2
=== RUN   TestPush/%!d(bool=true),3
--- PASS: TestPush (0.00s)
    --- PASS: TestPush/10,1 (0.00s)
    --- PASS: TestPush/%!d(string=Test),2 (0.00s)
    --- PASS: TestPush/%!d(bool=true),3 (0.00s)
=== RUN   TestPopWhenStactEmpty
--- PASS: TestPopWhenStactEmpty (0.00s)
=== RUN   TestPop
=== RUN   TestPop/Pop_0
=== RUN   TestPop/Pop_1
=== RUN   TestPop/Pop_2
--- PASS: TestPop (0.00s)
    --- PASS: TestPop/Pop_0 (0.00s)
    --- PASS: TestPop/Pop_1 (0.00s)
    --- PASS: TestPop/Pop_2 (0.00s)
=== RUN   TestIsEmpty
--- PASS: TestIsEmpty (0.00s)
=== RUN   TestTopWhenStackEmpty
--- PASS: TestTopWhenStackEmpty (0.00s)
=== RUN   TestTop
--- PASS: TestTop (0.00s)
=== RUN   TestGetSizeWhenStackEmpty
--- PASS: TestGetSizeWhenStackEmpty (0.00s)
=== RUN   TestGetSize
--- PASS: TestGetSize (0.00s)
PASS
ok      Stack   0.634s
*/
