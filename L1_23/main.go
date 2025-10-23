package main

func main() {

	sl := make([]int, 10)
	i := 5
	copy(sl[i:], sl[i+1:])
	sl = sl[:len(sl)-1]

}
