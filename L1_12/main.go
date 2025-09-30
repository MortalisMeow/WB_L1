package main

import "fmt"

func main() {

	massiv := []string{"cat", "cat", "dog", "cat", "tree"}

	kor := make(map[string]struct{})

	for _, elem := range massiv {
		kor[elem] = struct{}{}
	}

	fmt.Println(kor)

}
