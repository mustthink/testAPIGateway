package handlers

import "net/http"

func (app *application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/profile", app.getUser)
	mux.HandleFunc("/microservice/name", app.nameService)

	return mux
}
