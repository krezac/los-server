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
	URL       string    `db:"URL"`
	Active    bool      `db:"ACTIVE"`
	Created   time.Time `db:"CREATED_TS"`
}

// User represents records in users table
type User struct {
	ID             int64     `db:"ID"`
	Login          string    `db:"LOGIN"`
	Password       string    `db:"PASSWORD"`
	RoleCompetitor bool      `db:"ROLE_COMPETITOR"`
	RoleJudge      bool      `db:"ROLE_JUDGE"`
	RoleDirector   bool      `db:"ROLE_DIRECTOR"`
	RoleAdmin      bool      `db:"ROLE_ADMIN"`
	Active         bool      `db:"ACTIVE"`
	Created        time.Time `db:"CREATED_TS"`
}

type InvalidToken struct {
	ID      int64     `db:"ID"`
	UserId  string    `db:"USER_ID"`
	Token   string    `db:"TOKEN"`
	ValidTo time.Time `db:"VALID_TO"`
}

type Competition struct {
	ID           int64     `db:"ID"`
	Name         string    `db:"NAME"`
	EventDate    time.Time `db:"EVENT_DATE"`
	RangeID      int64     `db:"RANGE_ID"`
	CategoryID   int64     `db:"CATEGORY_ID"`
	TypeID       int64     `db:"TYPE_ID"`
	Active       bool      `db:"ACTIVE"`
	Created      time.Time `db:"CREATED_TS"`
	CategoryCode string    `db:"CATEGORY_CODE"`
	CategoryName string    `db:"CATEGORY_NAME"`
	TypeCode     string    `db:"TYPE_CODE"`
	TypeName     string    `db:"TYPE_NAME"`
	RangeName    string    `db:"RANGE_NAME"`
}
