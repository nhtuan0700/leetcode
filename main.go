// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// dict := make(map[int]int)
	
	wg := sync.WaitGroup{}
	
	a := []int{}
	// for i := 0; i < 100000; i++ {
		// a = append(a, i)
	// }
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// dict[i] = i
			// a[i] = i
			a = append(a, i)
		}(i)
	}
	wg.Wait()

	fmt.Println("Hello, 世界")
}
