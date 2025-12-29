// This file demonstrates Sourcegraph Precise Code Navigation across package boundaries.
//
// == Testing Cross-Package Precise Navigation in Sourcegraph ==
//
// 1. Go to Definition (cross-package type):
//    Jump from a type usage to its definition in another package.
//    "PrefixGreeter" on line 33 → Highlights definition in internal/greeter/greeter.go line 11
//
// 2. Go to Definition (cross-package method):
//    Navigate from a method call to its implementation in another package.
//    "Greet" on line 34 → Highlights definition in internal/greeter/greeter.go line 16
//
// 3. Go to Definition (cross-package function):
//    Jump to a standalone function defined in a different package.
//    "Sum" on line 44 → Highlights definition in pkg/mathutil/mathutil.go line 4
//
// 4. Find Implementations (interface):
//    Discover all types implementing an interface across the codebase.
//    "Greeter" in internal/greeter/greeter.go line 6 → Shows PrefixGreeter (line 11) and CountingGreeter (line 21)
//
// 5. Find Implementations (interface method):
//    See all concrete implementations of an interface method.
//    "Greet" in internal/greeter/greeter.go line 7 → Shows implementations on lines 16 and 27
//
// 6. Find References (cross-package):
//    Find all usages of a symbol from another package.
//    "Sum" in pkg/mathutil/mathutil.go → Shows references in main.go line 44 AND mathutil.go line 15 (used by Average)
//
// 7. Go to Definition (pointer vs value receiver):
//    Navigate to methods with different receiver types.
//    "Count" on line 40 → Highlights pointer receiver method in internal/greeter/greeter.go line 33

package main

import (
	"fmt"

	"example.com/go-precise-multipkg/internal/greeter"
	"example.com/go-precise-multipkg/pkg/mathutil"
)

func main() {
	// Use the greeter package - value receiver implementation
	simple := greeter.PrefixGreeter{Prefix: "Hello,"}
	fmt.Println(simple.Greet("multipkg world"))

	// Use a counting greeter - pointer receiver implementation
	counter := &greeter.CountingGreeter{Prefix: "Welcome,"}
	fmt.Println(counter.Greet("first visitor"))
	fmt.Println(counter.Greet("second visitor"))
	fmt.Printf("Total greetings: %d\n", counter.Count())

	// Use the mathutil package - cross-package function calls
	scores := []int{85, 92, 78, 95, 88}
	fmt.Printf("Sum: %d\n", mathutil.Sum(scores...))
	fmt.Printf("Average: %.2f\n", mathutil.Average(scores...))
	fmt.Printf("Max: %d\n", mathutil.Max(scores...))
}
