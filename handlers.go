package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rafaeljusto/cctldstats/config"
	"github.com/rafaeljusto/cctldstats/db"
	"github.com/rafaeljusto/cctldstats/protocol"
)

func registeredDomains(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received a request for registered domains from %s", r.RemoteAddr)

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Error parsing the remote address. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	remoteAddr := net.ParseIP(host)
	if remoteAddr == nil {
		log.Printf("Invalid host %s", host)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var grant bool
	for _, ip := range config.CCTLDStats.ACL {
		if remoteAddr.Equal(ip) {
			grant = true
			break
		}
	}

	if !grant {
		log.Printf("IP %s is not authorized", remoteAddr)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", config.CCTLDStats.DomainTableName)
	row := db.Connection.QueryRow(query)

	var rdr protocol.RegisteredDomainsResponse
	if err = row.Scan(&rdr.Number); err != nil {
		log.Printf("Error retrieving the information from the database. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(rdr)
	if err != nil {
		log.Printf("Error encoding the response into JSON. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
