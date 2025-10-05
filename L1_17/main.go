package main

import "fmt"

func binSearch(m []int, find int) int {
	low := 0
	high := len(m) - 1

	for low <= high {
		mid := low + (high-low)/2
		if m[mid] == find {
			return mid
		} else if m[mid] < find {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	someslice := []int{0, 1, 2, 4, 7, 8, 10, 15, 35, 40}

	find := binSearch(someslice, 4)
	fmt.Println(find)

}
