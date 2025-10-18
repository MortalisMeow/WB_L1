package main

import (
	"WB/WB_L1/L1_21/Legacy"
	"WB/WB_L1/L1_21/NewService"
)

func main() {
	someWork := [2]NewService.AdapterForLegacy{NewService.NewLegacyAdapter(&Legacy.Legacy{}), &NewService.NewWork{}}

	for _, adapter := range someWork {
		adapter.SendSome()
	}
}
