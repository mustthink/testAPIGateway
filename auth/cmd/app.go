package main

import (
	"auth/internal"
	"log"
	"net/http"
	"os"
)

var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	connStr := "host=172.16.1.4 user=postgres password=123456 dbname=users sslmode=disable"
	addr := "172.16.1.5:8081"

	//connStr := "user=postgres password=123456 dbname=users sslmode=disable"
	//addr := "localhost:8081"

	pg, err := auth.OpenDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pg.Close()

	ser := auth.NewService(errorLog, pg, &addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  ser.Routes(),
	}

	log.Println("Running a server on", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
