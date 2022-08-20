package main

import (
	"log"
	"net/http"
	"os"
	"user/internal"
)

var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	connStr := "host=172.16.1.4 user=postgres password=123456 dbname=users sslmode=disable"
	addr := "172.16.1.2:8082"
	s := "secret"

	pg, err := auth.OpenDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pg.Close()

	ser := auth.NewService(errorLog, pg, &addr, &s)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  ser.Routes(),
	}

	log.Println("Running a server on", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
