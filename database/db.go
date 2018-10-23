package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	rangesColumns       = "ID, NAME, LATITUDE, LONGITUDE, URL, ACTIVE, CREATED_TS"
	competitionsColumns = "c.ID, c.NAME, c.EVENT_DATE, c.RANGE_ID, c.CATEGORY_ID, c.TYPE_ID, c.ACTIVE, c.CREATED_TS, cc.CODE as CATEGORY_CODE, cc.NAME as CATEGORY_NAME, ct.CODE as TYPE_CODE, ct.NAME as TYPE_NAME, r.NAME as RANGE_NAME"
	competitionsFrom    = "competitions c JOIN competition_categories cc ON c.CATEGORY_ID=cc.ID JOIN competition_types ct ON c.TYPE_ID=ct.ID JOIN ranges r on c.RANGE_ID=r.id"
	competitionsOrderBy = "c.EVENT_DATE, c.NAME ASC"
	usersColumns        = "ID, LOGIN, PASSWORD, ROLE_COMPETITOR, ROLE_JUDGE, ROLE_DIRECTOR, ROLE_ADMIN, ACTIVE, CREATED_TS"
	invalidTokenColumns = "ID, USER_ID, TOKEN, VALID_TO"
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

func (db *Database) GetRanges(activeOnly bool) ([]Range, error) {
	ranges := []Range{}
	query := "SELECT " + rangesColumns + " FROM ranges"
	if activeOnly {
		query += " WHERE ACTIVE <> 0"
	}
	query += " ORDER BY name ASC"
	err := db.db.Select(&ranges, query)
	return ranges, err
}

func (db *Database) GetRangeByID(id int64) (*Range, error) {
	r := Range{}
	query := "SELECT " + rangesColumns + " FROM ranges WHERE id=?"
	err := db.db.Get(&r, query, id)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (db *Database) GetCompetitions(rangeID int64, activeOnly bool) ([]Competition, error) {
	competitions := []Competition{}
	query := "SELECT " + competitionsColumns + " FROM " + competitionsFrom + " WHERE c.RANGE_ID=?"
	if activeOnly {
		query += " AND c.ACTIVE <> 0"
	}
	query += " ORDER BY " + competitionsOrderBy
	err := db.db.Select(&competitions, query, rangeID)
	return competitions, err
}

func (db *Database) GetCompetitionByID(id int64) (*Competition, error) {
	c := Competition{}
	query := "SELECT " + competitionsColumns + " FROM " + competitionsFrom + " WHERE c.id=?"
	err := db.db.Get(&c, query, id)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (db *Database) GetUserByID(id int64) (*User, error) {
	u := User{}
	err := db.db.Get(&u, "SELECT "+usersColumns+" FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *Database) GetUserByLogin(login string, activeOly bool) (*User, error) {
	u := User{}
	query := "SELECT " + usersColumns + " FROM users WHERE login=?"
	if activeOly {
		query += " AND ACTIVE<>0"
	}
	err := db.db.Get(&u, query, login)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *Database) InvalidateToken(login, token string, validTo time.Time) error {
	user, err := db.GetUserByLogin(login, true)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO invalid_tokens(USER_ID, TOKEN, VALID_TO) VALUES (?, ?, ?)", user.ID, token, validTo)
	return err
}

func (db *Database) IsTokenInvalid(login, token string) bool {
	user, err := db.GetUserByLogin(login, true)
	if err != nil {
		return true // user not found, so invalid
	}

	it := InvalidToken{}
	err = db.db.Get(&it, "SELECT "+invalidTokenColumns+" FROM invalid_tokens where USER_ID=? AND TOKEN=? LIMIT 1", user.ID, token)
	return err == nil // means the read succeeded
}
