package repository

import (
	"github.com/y-kouhei9/bookings-app/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
