package infrastructures

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDatabase() (db *sql.DB, err error) {
	connStr := "user=username dbname=inventorydb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
