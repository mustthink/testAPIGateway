package main

import (
	"log"
	"net/http"
	"os"
	"proxy/internal/handlers"
)

func main() {
	//connStr := "host=172.16.1.4 user=postgres password=123456 dbname=users sslmode=disable"
	addr := "localhost:8083"

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := handlers.NewApplication(errorLog, &addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	log.Println("Запуск веб-сервера на", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}