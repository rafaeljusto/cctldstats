package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rafaeljusto/cctldstats/config"
	"github.com/rafaeljusto/cctldstats/db"
)

func registeredDomains(w http.ResponseWriter, r *http.Request) {
	remoteAddr := net.ParseIP(r.RemoteAddr)
	if remoteAddr == nil {
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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", config.CCTLDStats.DomainTableName)
	row := db.Connection.QueryRow(query)

	var rdr RegisteredDomainsResponse
	if err := row.Scan(&rdr.Number); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(rdr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
