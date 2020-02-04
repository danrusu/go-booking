package main

func main() {
	println("Simple FOR")
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}

	println("FOR POST CLAUSE")
	for j := 0; j < 5; j++ {
		println(j)
	}

	println("Infinite")
	var k int
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}

	println("For Collections")
	slice := []int{1, 2, 3}

	println("For Collections classic")
	for ii := 0; ii < len(slice); ii++ {
		println(slice[ii])
	}

	println("For slice")
	for i, v := range slice {
		println(i, v)
	}

	println("For map")
	m := map[string]int{"one": 1, "two": 2, "ten": 10}
	for k, v := range m {
		println(k, v)
	}

	println("For map keys")
	for k := range m {
		println(k)
	}

	println("For map values")
	for _, v := range m {
		println(v)
	}
}
