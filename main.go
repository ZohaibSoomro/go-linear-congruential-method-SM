package main

import (
	"fmt"
	"strconv"
)

func main() {
	X := 7
	a := 5
	c := 3
	m := 16
	randomIntegers := make([]int, 10)
	randomNumbers := make([]float64, 10)

	fmt.Print("\nXi+1= (a.Xi+c) mod m  :-i=0,1,2...\n")
	fmt.Printf("Ri= Xi/m    :-i=1,2,3...\n\n")
	for i := range randomIntegers {
		// oldX := X
		multAndAddition := (a*X + c)
		X = multAndAddition % m
		randomIntegers[i] = X
		myFloat := float64(X) / float64(m)
		str := fmt.Sprintf("%.2f", myFloat)
		randomNumbers[i], _ = strconv.ParseFloat(str, 64)
		// fmt.Println("X", i+1, "= (", a, "x", oldX, "+", c, ") mod", m, "=", multAndAddition, "mod", m, "=", X)
		// fmt.Println("R", i+1, "=", X, "/", m, " = ", randomNumbers[i])
	}
	// fmt.Printf("\n%-10s%-20s%-20s\n", "i", "Xi", "Ri")
	// for i := range randomIntegers {
	// 	fmt.Printf("%-10d%-20d%-20f\n", i+1, randomIntegers[i], randomNumbers[i])
	// }

	fmt.Println("Random Integers: ", randomIntegers)
	fmt.Println("Random Numbers: ", randomNumbers)
}
