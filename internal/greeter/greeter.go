package greeter

import "fmt"

// Greeter defines the greeting interface.
type Greeter interface {
	Greet(name string) string
}

// PrefixGreeter implements Greeter with a configurable prefix.
type PrefixGreeter struct {
	Prefix string
}

// Greet returns a prefixed greeting for the given name.
func (g PrefixGreeter) Greet(name string) string {
	return g.Prefix + " " + name
}

// CountingGreeter tracks how many greetings have been made.
type CountingGreeter struct {
	Prefix string
	count  int
}

// Greet returns a greeting and increments the internal counter.
func (g *CountingGreeter) Greet(name string) string {
	g.count++
	return fmt.Sprintf("[#%d] %s %s", g.count, g.Prefix, name)
}

// Count returns the number of greetings made.
func (g *CountingGreeter) Count() int {
	return g.count
}
