package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Выход по условию
//через канал уведомления
//через контекст
//прекращение работы runtime.Goexit()

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	//Exit with condition
	var condition = true
	wg.Add(1)
	go func() {
		defer wg.Done()
		if condition == true {
			fmt.Println("Exit with condition")
			return
		}
	}()

	//Exit with channel
	notif := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-notif:
				fmt.Println("Exit with channel")
				return
			default:
				fmt.Println("Some work for channel ")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	notif <- true

	//Exit with Context
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exit with context")
				return
			default:
				fmt.Println("Some work for context example")
				time.Sleep(1 * time.Second)
			}
		}

	}()
	cancel()

	//Exit with runtime.Goexit()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Exit with Goexit")
		runtime.Goexit()
	}()

	wg.Wait()
}
