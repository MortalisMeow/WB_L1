package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, tasks chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение работы воркера")
			return
		case res, ok := <-tasks:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %d: %d\n", id, res)

		}
	}
}
func main() {
	var n int
	result := 0
	tasks := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Введите количество воркеров:")
	fmt.Scan(&n)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
		close(tasks)
	}()

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, tasks, &wg)
	}

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return

		case tasks <- result:
			result++
		}
	}

}
