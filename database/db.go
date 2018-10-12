package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database is wrapper around database connection
type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sql.DB, driverName string) *Database {
	database := Database{}
	database.db = sqlx.NewDb(db, driverName)
	return &database
}

func NewMysqlDatabase() (*Database, error) {
	database := Database{}
	db, err := sql.Open("mysql", "los:los@/los?parseTime=true")
	if err != nil {
		return nil, err
	}
	database.db = sqlx.NewDb(db, "mysql")
	return &database, nil
}

func (db *Database) GetRanges() ([]Range, error) {
	ranges := []Range{}
	err := db.db.Select(&ranges, "SELECT * FROM ranges ORDER BY name ASC")
	return ranges, err
}

func (db *Database) GetRangeByID(id int64) (*Range, error) {
	r := Range{}
	err := db.db.Get(&r, "SELECT * FROM ranges WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
