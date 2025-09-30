package main

import "fmt"

func main() {

	A := 2
	B := 7

	//XOR обмен
	A ^= B

	B ^= A

	A ^= B

	fmt.Println(A, B)

}
