package main

import (
	"fmt"
	"stockExchangeSim/investor"
)

func main() {
	firstInvestor := investor.NewInvestor("Fred")
	fmt.Printf("First investor name: %s", firstInvestor.GetName())
}
