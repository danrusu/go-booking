package main

import (
	"net/http"
)

func main() {

	controllers.registerControllers()
	http.ListenAndServe(":1111", nil)

}
