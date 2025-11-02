package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println(time.Now())
	Sleep(5 * time.Second)
	fmt.Println(time.Now())
}
