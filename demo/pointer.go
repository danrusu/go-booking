package main

import (
	"fmt"
)

func main() {
	var name *string = new(string)
	*name = "Adela"
	fmt.Println(name)
	fmt.Println(*name) // dereferentiate

	name2 := "Tania"
	ptr := &name2
	fmt.Println(ptr, *ptr)

	name2 = "Maria"
	fmt.Println(ptr, *ptr)
}
