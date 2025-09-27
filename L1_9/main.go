package main

import (
	"fmt"
	"sync"
)

func main() {
	massiv := [4]int{1, 2, 3, 4}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for _, num := range massiv {
			ch1 <- num
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch2)
		for res := range ch1 {
			ch2 <- res * 2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range ch2 {
			fmt.Println(res)
		}
	}()

	wg.Wait()
}
