package main

import "fmt"

func optype(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

func main() {

	a := "a"

	b := 1

	c := true

	ch := make(chan int)

	optype(a)
	optype(b)
	optype(c)
	optype(ch)

}
