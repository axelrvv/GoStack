package main

import "fmt"

var arr []int

func main() {
	arr = append(arr, 1)
	arr = append(arr, 4)

	lengh := len(arr)

	for index, x := range arr {
		fmt.Println(index, x)
	}
	fmt.Println(lengh)

}
