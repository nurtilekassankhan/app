package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Printer
	Textile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Printer: NewPrinterSQLite(db),
		Textile: NewTextileSQLite(db),
	}
}
