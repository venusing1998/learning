package main

import (
	"fmt"
)

func main() {
	a := []float64{67.7, 89.8, 21, 78}
	fmt.Println(a)
	sum := float64(0)
	for i, v := range a { // range returns both the index and value
		fmt.Printf("%d the elemrnt of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elemrnts of a", sum)
}
