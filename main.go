package main

import (
	"net/http"

	"github.com/danrusu/learn-go-webservice/booking/controllers"
)

func main() {

	controllers.registerControllers()
	http.ListenAndServe(":1111", nil)

}
