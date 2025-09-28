package main

import "fmt"

func main() {

	A := [3]int{1, 2, 3}
	B := [3]int{2, 3, 4}

	kor := make(map[int]struct{})

	for _, a := range A {
		for _, b := range B {
			if a == b {
				kor[a] = struct{}{}
			}
		}
	}

	fmt.Println(kor)

}
