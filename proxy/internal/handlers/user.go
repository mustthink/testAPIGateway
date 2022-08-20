package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	uname := r.URL.Query().Get("username")
	if uname == "" {
		w.Header().Set("Username", "not found")
		app.clientError(w, http.StatusNotFound)
		return
	}

	req, err := http.NewRequest("GET", "http://172.16.1.5:8081/auth", nil)
	req.Header.Set("Username", uname)

	resp, err := app.client.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		w.Header().Set("Username", "not found")
		app.clientError(w, http.StatusUnauthorized)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	req, err = http.NewRequest("GET", "http://172.16.1.2:8082/user/profile", nil)
	req.Header.Set("Username", uname)
	resp, err = app.client.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsondata, err := io.ReadAll(resp.Body)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	resp.Body.Close()

	fmt.Fprintf(w, "%v", string(jsondata))
}

func (app *application) nameService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	resp, err := app.client.Get("http://172.16.1.2:8082/microservice/name")
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Fprintf(w, "%v", string(data))
}
