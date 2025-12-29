package main

import (
	"fmt"

	"example.com/go-precise-multipkg/internal/greeter"
	"example.com/go-precise-multipkg/pkg/mathutil"
)

func main() {
	// Use the greeter package
	simple := greeter.PrefixGreeter{Prefix: "Hello,"}
	fmt.Println(simple.Greet("multipkg world"))

	// Use a counting greeter to demonstrate method with receiver pointer
	counter := &greeter.CountingGreeter{Prefix: "Welcome,"}
	fmt.Println(counter.Greet("first visitor"))
	fmt.Println(counter.Greet("second visitor"))
	fmt.Printf("Total greetings: %d\n", counter.Count())

	// Use the mathutil package - demonstrates cross-package symbol resolution
	scores := []int{85, 92, 78, 95, 88}
	fmt.Printf("Sum: %d\n", mathutil.Sum(scores...))
	fmt.Printf("Average: %.2f\n", mathutil.Average(scores...))
	fmt.Printf("Max: %d\n", mathutil.Max(scores...))
}
