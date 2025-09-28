package main

import (
	"fmt"
	"sync"
)

func sort(temp float64, tempMap map[int][]float64) {
	var mu sync.Mutex

	key := (int(temp) / 10) * 10
	mu.Lock()
	tempMap[key] = append(tempMap[key], temp)
	mu.Unlock()
}

func main() {

	posl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -29.9}
	tempGroup := make(map[int][]float64)

	for num := range posl {
		sort(posl[num], tempGroup)
	}

	for key, num := range tempGroup {
		fmt.Printf("%d: %v\n", key, num)
	}

}
