package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/danrusu/go-booking/booking/models"
)

type bookingController struct {
	BookingIDPattern *regexp.Regexp
}

func (controller bookingController) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/bookings" {
		switch request.Method {
		case http.MethodGet:
			controller.getAll(responseWriter, request)
		case http.MethodPost:
			controller.post(responseWriter, request)
		default:
			responseWriter.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := controller.BookingIDPattern.FindStringSubmatch(request.URL.Path)
		if len(matches) == 0 {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}
		switch request.Method {
		case http.MethodGet:
			controller.get(id, responseWriter)
		case http.MethodPut:
			controller.put(id, responseWriter, request)
		case http.MethodDelete:
			controller.delete(id, responseWriter)
		default:
			responseWriter.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (controller *bookingController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetBookings(), w)
}

func (controller *bookingController) get(id int, w http.ResponseWriter) {
	u, err := models.GetBooking(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (controller *bookingController) post(w http.ResponseWriter, r *http.Request) {
	u, err := controller.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Booking object"))
		return
	}
	u, err = models.CreateBooking(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (controller *bookingController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := controller.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Booking object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted Booking must match ID in URL"))
		return
	}
	u, err = models.UpdateBooking(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (controller *bookingController) delete(id int, w http.ResponseWriter) {
	err := models.DeleteBooking(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (controller *bookingController) parseRequest(r *http.Request) (models.Booking, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Booking
	err := dec.Decode(&u)
	if err != nil {
		return models.Booking{}, err
	}
	return u, nil
}

func newBookingController() *bookingController {
	return &bookingController{
		BookingIDPattern: regexp.MustCompile(`^/bookings/(\d+)/?`),
	}
}
