package database

import (
	"time"
)

// Range represents records in ranges table
type Range struct {
	ID        int64     `db:"ID"`
	Name      string    `db:"NAME"`
	Latitude  float64   `db:"LATITUDE"`
	Longitude float64   `db:"LONGITUDE"`
	Created   time.Time `db:"CREATED_TS"`
	Active    bool      `db:"ACTIVE"`
}
