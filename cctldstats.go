package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/rafaeljusto/cctldstats/config"
	"github.com/rafaeljusto/cctldstats/db"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/domains/registered", http.HandlerFunc(registeredDomains))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
