package main

import (
	"net/http"

	"github.com/danrusu/go-booking/booking/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":1111", nil)

}
