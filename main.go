package main

import (
	"fmt"
	"math"
)

func main() {
  // fmt.Printf("Привет")
  
	var userHeight = 1.8
	var userWeight float64 = 92
	var IMT = userWeight / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
