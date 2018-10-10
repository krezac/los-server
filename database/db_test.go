package database_test

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/krezac/los-server/database"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

/*
func newMockDatabase() (*database.Database, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	// defer mockDB.Close()
	database := database.NewDatabase(mockDB, "sqlmock")
	return database, mock, nil
}
*/
var _ = Describe("Database", func() {
	var mockDB *sql.DB
	var db *database.Database
	var mock sqlmock.Sqlmock
	var mockErr error
	//var err error

	//BeforeEach(func() {
	//
	//})

	Describe("Testing db wrapper", func() {
		Context("for ranges", func() {

			columns := []string{"ID", "NAME", "LATITUDE", "LONGITUDE", "ACTIVE"}

			//mock.ExpectBegin()
			// expect query to fetch order and user, match it with regexp
			mockDB, mock, _ = sqlmock.New()
			mock.ExpectQuery("SELECT \\* FROM ranges ORDER BY name ASC").
				WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Strelnice 1", 1, 2, 1).AddRow(2, "Strelnice 2", 3, 4, 1).AddRow(3, "Strelnice 3", 5, 6, 1))
			db = database.NewDatabase(mockDB, "sqlmock")
			// expect transaction rollback, since order status is "cancelled"
			//mock.ExpectRollback()

			Context("when retrieving list", func() {

				var ranges []database.Range
				var err error
				ranges, err = db.GetRanges()

				It("should return three records", func() {
					Expect(db).NotTo(BeNil())
					Expect(mock).NotTo(BeNil())
					Expect(mockErr).To(BeNil())

					Expect(len(ranges)).To(Equal(3))
					Expect(err).To(BeNil())
				})

			})

		})
	})

})

/*
func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
*/
