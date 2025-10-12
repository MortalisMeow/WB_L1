package main

import (
	"fmt"
)

func reversew(s string) string {
	b := []byte(s)

	rev := func(b []byte, i, j int) {
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	rev(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			rev(b, start, i-1)
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	s := "snow dog sun"
	fmt.Println(reversew(s))
}
