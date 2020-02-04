package main

type User struct {
	ID   int
	Name string
}

func main() {
	u1 := User{
		ID:   1,
		Name: "dan",
	}

	u2 := User{
		ID:   2,
		Name: "tania",
	}

	if u1 == u2 {
		println("Same user!")
	} else if u1.Name == u2.Name {
		println("Similar users!")
	} else {
		println("Different users!")
	}
}
