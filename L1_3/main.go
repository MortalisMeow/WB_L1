package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for res := range tasks {
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d: %d\n", id, res)
	}

}
func main() {
	var n int
	result := 0
	var wg sync.WaitGroup

	fmt.Println("Введите количество воркеров:")
	fmt.Scan(&n)

	tasks := make(chan int)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for {
		tasks <- result
		result++
	}

	wg.Wait()
}
