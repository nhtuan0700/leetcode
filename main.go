package main

import "fmt"

func check(a []int) {
	a = append(a, 1)
}

func main() {
	a := make([]int, 1)
	fmt.Println(a)

	check(a)
	fmt.Println(a)
}
