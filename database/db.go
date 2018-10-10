package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// Database is wrapper around database connection
type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sql.DB, driverName string) *Database {
	database := Database{}
	database.db = sqlx.NewDb(db, driverName) // sqlx.Connect("mysql", "root:root@/los")
	return &database
}

func (db *Database) GetRanges() ([]Range, error) {
	ranges := []Range{}
	err := db.db.Select(&ranges, "SELECT * FROM ranges ORDER BY name ASC")
	return ranges, err
}
