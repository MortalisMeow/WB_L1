package main

import (
	"fmt"
	"sync"
)

type MuMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewMuMap() *MuMap {
	return &MuMap{
		m: make(map[string]int),
	}
}

func main() {
	var wg sync.WaitGroup

	//Использование мьютекса
	m := NewMuMap()

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.mu.Lock()
		m.m["first"] = 1
		fmt.Println("Mutex goroutine done")
		m.mu.Unlock()
	}()

	//Использование потокобезопасной мапы
	var syncM sync.Map
	wg.Add(1)
	go func() {
		defer wg.Done()
		syncM.Store("first", 1)
		fmt.Println("sync.Map goroutine done")
	}()

	wg.Wait()
}
