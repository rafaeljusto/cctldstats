package db

import (
	"database/sql"
	"fmt"

	"github.com/rafaeljusto/cctldstats/config"
)

// Connection database connection.
var Connection *sql.DB

// Connect performs the database connection. Today the following databases are supported: mysql and postgres
func Connect() (err error) {
	var connParams string
	switch config.CCTLDStats.Database.Kind {
	case "mysql":
		connParams = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			config.CCTLDStats.Database.Username,
			config.CCTLDStats.Database.Password,
			config.CCTLDStats.Database.Host,
			config.CCTLDStats.Database.Name,
		)
	case "postgres":
		connParams = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			config.CCTLDStats.Database.Username,
			config.CCTLDStats.Database.Password,
			config.CCTLDStats.Database.Host,
			config.CCTLDStats.Database.Name,
		)
	}

	Connection, err = sql.Open(config.CCTLDStats.Database.Kind, connParams)
	return
}
