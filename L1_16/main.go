package main

import "fmt"

func quickSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	arr := make([]int, len(m))
	copy(arr, m)

	left := 0
	right := len(arr) - 1
	mid := arr[(left+right)/2]

	i := left
	j := right

	for i <= j {
		for arr[i] < mid {
			i++
		}
		for arr[j] > mid {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	middle := arr[j+1 : i]
	leftSort := quickSort(arr[:j+1])
	rightSort := quickSort(arr[i:])

	result := append(leftSort, middle...)
	result = append(result, rightSort...)

	return result
}

func main() {

	someSlice := []int{1, 5, 20, 46, 4, 28, 0, 0, 35, 3}

	result := quickSort(someSlice)

	fmt.Println(result)

}
