package config

import (
	"net"

	"github.com/kelseyhightower/envconfig"
)

const prefix = "cctldstats"

// CCTLDStats global project configuration.
var CCTLDStats *cctldstatsConfig

// config contains all configuration parameters for running the statistic services.
type cctldstatsConfig struct {
	// Database stores all the database information to connect to the back-end.
	Database struct {
		// King type of the database, possible values are: mysql or postgres.
		Kind string `envconfig:"kind"`

		// Name name of the database.
		Name string `envconfig:"name"`

		// Username user used to connect to the desired database.
		Username string `envconfig:"username"`

		// Password password used to connect to the desired database.
		Password string `envconfig:"password"`

		// Host address of the database
		Host string `envconfig:"host"`
	} `envconfig:"database"`

	// DomainTableName name of the table in the database that stores the domains
	DomainTableName string `envconfig:"domain_table_name"`

	// ACL IP addresses that are allowed to retrieve information from the services.
	ACL []net.IP `envconfig:"acl"`
}

// Load fill the global configuration variable using default values and environment variables.
func Load() error {
	CCTLDStats = new(cctldstatsConfig)
	CCTLDStats.Database.Kind = "mysql"
	CCTLDStats.Database.Host = "localhost"
	CCTLDStats.DomainTableName = "domain"
	CCTLDStats.ACL = []net.IP{net.ParseIP("127.0.0.1")}
	return envconfig.Process(prefix, CCTLDStats)
}
