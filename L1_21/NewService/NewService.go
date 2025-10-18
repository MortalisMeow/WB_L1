package NewService

import (
	"WB/WB_L1/L1_21/Legacy"
	"fmt"
)

type AdapterForLegacy interface {
	SendSome()
}

type NewWork struct {
}

func (n *NewWork) SendSome() {
	fmt.Println("From new service")
}

type LegacyAdapter struct {
	*Legacy.Legacy
}

func NewLegacyAdapter(new *Legacy.Legacy) AdapterForLegacy {
	return &LegacyAdapter{new}
}

func (l *LegacyAdapter) SendSome() {
	fmt.Println("From Legacy, with new message")

}
