package database_test

import (
	"testing"

	"github.com/krezac/los-server/database"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var rangesColumns = []string{"ID", "NAME", "LATITUDE", "LONGITUDE", "ACTIVE"}

func TestGetRanges(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NotNil(t, mockDB, "mock created")
	assert.NotNil(t, mock, "mock created")
	assert.NoError(t, err, "mock created")

	mock.ExpectQuery("SELECT \\* FROM ranges ORDER BY name ASC").
		WillReturnRows(sqlmock.NewRows(rangesColumns).AddRow(
			11, "Strelnice 1", 1, 2, 1).AddRow(
			12, "Strelnice 2", 3, 4, 1).AddRow(
			13, "Strelnice 3", 5, 6, 1))

	db := database.NewDatabase(mockDB, "sqlmock")
	assert.NotNil(t, mockDB, "database wrapper created")

	ranges, err := db.GetRanges()
	assert.Len(t, ranges, 3, "list size")
	assert.Equal(t, 11, ranges[0].ID, "data check")
	assert.Equal(t, 12, ranges[1].ID, "data check")
	assert.Equal(t, 13, ranges[2].ID, "data check")
}

/*
func TestGetRangesDb(t *testing.T) {

	// note this test requires real database
	db, err := database.NewMysqlDatabase()
	assert.NotNil(t, db, "database wrapper created")
	assert.NoError(t, err)

	ranges, err := db.GetRanges()
	assert.Len(t, ranges, 2, "list size")
	assert.Equal(t, 1, ranges[0].ID, "data check")
	assert.Equal(t, 2, ranges[1].ID, "data check")
}
*/
