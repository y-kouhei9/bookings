package helpers

import (
	"runtime/debug"
	"fmt"
	"net/http"

	"github.com/y-kouhei9/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers.
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError 
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError 
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
