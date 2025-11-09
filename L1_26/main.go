package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	uMap := make(map[rune]bool)
	for _, a := range strings.ToLower(str) {
		if uMap[a] == false {
			uMap[a] = true
		} else if uMap[a] == true {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}
