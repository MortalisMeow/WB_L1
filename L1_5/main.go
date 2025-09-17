package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.NewTimer(3 * time.Second)
	defer time.Stop()
	ch := make(chan int)
	val := 1

	go func() {
		for {
			select {
			case <-time.C:
				return
			case res, ok := <-ch:
				if !ok {
					time.Stop()
					return
				}
				fmt.Printf("result : %d\n", res)

			}
		}

	}()

	for {
		select {
		case <-time.C:
			return
		case ch <- val:
			val++
		}

	}

}
