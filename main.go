package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zuzannatomaszyk/goApi/db"
	"github.com/zuzannatomaszyk/goApi/handler"
)

func main() {
	addr := ":8081"

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handler.HandleRequests(database)
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, httpHandler))
}
