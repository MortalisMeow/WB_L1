package Legacy

import "fmt"

type Legacy struct{}

func (l *Legacy) DoSomeText() {
	fmt.Println("From Legacy")
}
