package main

import "fmt"

func main() {

	var input string
	fmt.Printf("Введите строку: ")
	fmt.Scan(&input)

	str := []rune(input)
	newstr := make([]rune, len(str))

	for i := 0; i < len(str); i++ {
		newstr[(len(str)-1)-i] = str[i]
	}
	fmt.Println(string(newstr))
}
