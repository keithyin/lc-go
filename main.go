package main

import "fmt"

func main() {
	fmt.Println("hello world")

	a := []int{2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}

	a = make([]int, 10)
	fmt.Println(a)
}
