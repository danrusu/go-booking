package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARRAYS")
	// long form
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println(arr1)

	// short form
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	fmt.Println("SLICES")
	slice1 := arr2[:] // similar to a pointer to arr2
	fmt.Println(slice1)

	// changes in array are reflected in the slice
	arr2[1] = 100
	slice2 := arr2[1:2]
	fmt.Println(arr2, slice2)

	// changes in slice are reflected in the array
	slice2[0] = 200
	fmt.Println(arr2, slice1)

	// better slice, underline array managed by GO
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4, 5, 100, 200)
	fmt.Println(slice)

	fmt.Println("MAPS")
	m := map[string]int{"one": 1, "two": 2, "ten": 10}
	fmt.Println(m)
	fmt.Println(m["ten"])
	delete(m, "two")
	fmt.Println(m)
}
