package handlers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type application struct {
	errorLog *log.Logger
	url      *string
	client   *http.Client
}

func NewApplication(errorLog *log.Logger, url *string) *application {
	return &application{
		errorLog: errorLog,
		url:      url,
		client:   &http.Client{},
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
