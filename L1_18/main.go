package main

import (
	"fmt"
	"sync"
)

type concurStruct struct {
	mu sync.Mutex
	c  int
}

func (c *concurStruct) inc() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	newCon := &concurStruct{c: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newCon.inc()
		}()
	}

	wg.Wait()
	fmt.Println(newCon.c)

}
