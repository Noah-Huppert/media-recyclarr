package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDBOpts are options for OpenDB()
type OpenDBOpts struct {
	// PostgresURI is a connection URI for Postgres
	PostgresURI string
}

// OpenDB connects to the database
func OpenDB(opts OpenDBOpts) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opts.PostgresURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	driverDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve underlying driver database client: %s", err)
	}
	if err := driverDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %s", err)
	}

	return db, nil
}
