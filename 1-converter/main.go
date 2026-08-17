package main

import "fmt"

func main() {
	const USDtoEUR float64 = 0.86
	const USDtoRUR float64 = 85
	EURtoRUR := USDtoRUR / USDtoEUR
	fmt.Print(EURtoRUR)
}
