package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	ordersTable   = "orders"
	shirtsTable   = "shirts"
	paintsTable   = "paints"
	papersTable   = "papers"
	textilesTable = "textiles"
	printersTable = "printers"
)

func NewSQLiteDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./app.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
