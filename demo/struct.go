package main

import (
	"fmt"
)

func main() {

	type user struct {
		ID   int
		Name string
	}
	var u user
	u.ID = 1000
	u.Name = "Dorina"

	fmt.Println(u)

	// short initializer
	u2 := user{
		ID:   1001,
		Name: "Dan",
	}
	fmt.Println(u2)
}
