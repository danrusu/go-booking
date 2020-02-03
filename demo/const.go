package main

import (
	"fmt"
)

// iota
const (
	first  = iota + 6  // iota = 0
	second = 2 << iota // iota = 1
	third              // will repeat previous 2 << iota

	pi = 3.1415 // determined at compile time
)

func main() {
	// cannot return from a function, functions are evaluated at runtime
	fmt.Println(pi, first, second, third)
}
