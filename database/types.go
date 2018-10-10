package database

// Range represents records in ranges table
type Range struct {
	ID        int     `db:"ID"`
	Name      string  `db:"NAME"`
	Latitude  float32 `db:"LATITUDE"`
	Longitude float32 `db:"LONGITUDE"`
	Active    bool    `db:"ACTIVE"`
}
