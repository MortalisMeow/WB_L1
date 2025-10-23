package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1000000000000000000", 10)
	b.SetString("5000000000000000000", 10)

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())

	result := new(big.Int)

	result.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", result.String())

	result.Add(a, b)
	fmt.Printf("Сложение: %s\n", result.String())

	result.Div(a, b)
	fmt.Printf("Деление: %s\n", result.String())

	result.Mul(a, b)
	fmt.Printf("Умножение: %s\n", result.String())

	result.Mod(a, b)
	fmt.Printf("Остаток от деления: %s\n", result.String())
}
