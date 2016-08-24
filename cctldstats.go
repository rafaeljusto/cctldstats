package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/http"
	"strconv"
)

var (
	database struct {
		kind     string // mysql or postgres
		name     string
		username string
		password string
		host     string
	}

	domainTable string

	acl []net.IP
)

func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/domains/registered", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := net.ParseIP(r.RemoteAddr)
		if remoteAddr == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var grant bool
		for _, ip := range acl {
			if remoteAddr.Equal(ip) {
				grant = true
				break
			}
		}

		if !grant {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		query := fmt.Sprintf("SELECT COUNT(*) FROM %s", domainTable)

		row := db.QueryRow(query)

		var registeredDomains int
		if err := row.Scan(&registeredDomains); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(strconv.Itoa(registeredDomains)))
	}))

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func connectDatabase() (*sql.DB, error) {
	var connParams string
	switch database.kind {
	case "mysql":
		connParams = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			database.username,
			database.password,
			database.host,
			database.name,
		)
	case "postgres":
		connParams = fmt.Sprintf("postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full",
			database.username,
			database.password,
			database.host,
			database.name,
		)
	}

	return sql.Open(database.kind, connParams)
}
