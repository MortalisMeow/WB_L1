package main

import "fmt"

func toZero(n int64, i int64) int64 {
	return n &^ (1 << i)
}

func toOne(n int64, i int64) int64 {
	return n | (1 << i)
}

func main() {

	var n int64
	var i int64

	fmt.Println("Введите число:")
	fmt.Scan(&n)

	fmt.Println("Введите i бит:")
	fmt.Scan(&i)
	if i < 0 || i > 63 {
		fmt.Println("Ошибка: бит должен быть между 0 и 63")
		return
	}

	zero := toZero(n, i)
	one := toOne(n, i)

	//i бит в 0
	fmt.Printf("Result for i bit to 0 : %v (%08b)\n", zero, zero)

	//i бит в 1
	fmt.Printf("Result for i bit to 1 : %v (%08b)\n", one, one)
}
