package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func registerControllers() {
	controller := newBookingController()

	http.Handle("/bookings", *controller)
	http.Handle("/bookings/", *controller)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(data)
}
