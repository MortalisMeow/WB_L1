package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	list := [5]int{2, 4, 6, 8, 10}

	for index, elem := range list {
		wg.Add(1)
		go func(index int, elem int) {
			defer wg.Done()
			list[index] = elem * elem
		}(index, elem)
	}

	wg.Wait()

	for _, elem := range list {
		fmt.Println(elem)
	}
}
