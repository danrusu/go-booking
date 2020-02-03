package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port, 2) // _ is the write only variable
	fmt.Println(err)

}

//func startWebServer(port int, numberOfRetries int) {
func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")

	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, errors.New("Something went wrong")
}
