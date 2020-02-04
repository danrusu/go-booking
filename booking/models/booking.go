package models

import (
	"errors"
	"fmt"
)

type Booking struct {
	ID              int
	FirstName       string
	LastName        string
	Price           float32
	DepositPaid     bool
	Checkin         string
	Checkout        string
	AdditionalNeeds string
}

var (
	bookings      []*Booking
	nextBookingID = 1
)

func GetBookings() []*Booking {
	return bookings
}

func CreateBooking(booking Booking) (Booking, error) {
	if booking.ID != 0 {
		return Booking{}, errors.New("New Booking must not include id or it must be set to zero")
	}
	booking.ID = nextBookingID
	nextBookingID++
	bookings = append(bookings, &booking)
	return booking, nil
}

func GetBooking(id int) (Booking, error) {
	for _, booking := range bookings {
		if booking.ID == id {
			return *booking, nil
		}
	}

	return Booking{}, fmt.Errorf("Booking with ID '%v' not found", id)
}

func UpdateBooking(booking Booking) (Booking, error) {
	for index, updateCandidate := range bookings {
		if updateCandidate.ID == booking.ID {
			bookings[index] = &booking
			return booking, nil
		}
	}

	return Booking{}, fmt.Errorf("Booking with ID '%v' does not exist", booking.ID)
}

func removeBookingFromIndex(bookings []*Booking, index int) {
	bookings = append(bookings[:index], bookings[index+1:]...)
}

func DeleteBooking(id int) error {
	for index, booking := range bookings {
		if booking.ID == id {
			removeBookingFromIndex(bookings, index)
			return nil
		}
	}

	return fmt.Errorf("Booking with ID '%v' does not exist", id)
}
