package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.After(3 * time.Second)
	ch := make(chan int)
	val := 1

	go func() {
		for {
			select {
			case <-time:
				return
			case res := <-ch:
				fmt.Printf("result : %d\n", res)

			}
		}

	}()

	for {
		select {
		case <-time:
			return
		case ch <- val:
			val++
		}

	}

}
