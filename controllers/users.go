package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPAttern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))

}

func newUserController() *userController {
	return &userController{
		userIDPAttern: regexp.MustCompile(`^/users/(\d+)/?`),
	}

}
