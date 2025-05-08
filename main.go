package main

import "fmt"

func check(a []int) {
	a = append(a, 1)
}

func main() {
	for i := range 4 {
		fmt.Println(i)
	}
}
