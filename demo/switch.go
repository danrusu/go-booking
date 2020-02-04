package main

type HTTPRequest struct {
	Method string
}

func main() {
	//r := HTTPRequest{Method: "GET"}
	r := HTTPRequest{Method: "DELETE"}

	switch r.Method {

	case "GET":
		println("Get request")
		fallthrough
	case "POST":
		println("POST request") // break by default

	case "PUT":
		println("Put request")

	default:
		println("Unhandled method")
	}

}
