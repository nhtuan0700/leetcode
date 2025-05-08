package main

import "fmt"

func main() {
	a := []int{0}

	// b := a[:]
	fmt.Println(len(a)) // 10
	a = a[1:]
	fmt.Println(len(a)) // 10

	
	// c[4] = 7

	// fmt.Printf("%p\n", &b[9])
	// fmt.Printf("%p\n", &a[9])
	// fmt.Printf("%p\n", &c[4])
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
}
