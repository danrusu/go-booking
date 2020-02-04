package main

import (
	"net/http"

	"github.com/danrusu/learn-go-webservice/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
