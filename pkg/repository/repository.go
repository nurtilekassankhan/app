package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Printer
	Textile
	Paper
	Paint
	Shirt
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Printer: NewPrinterSQLite(db),
		Textile: NewTextileSQLite(db),
		Paper:   NewPaperSQLite(db),
		Paint:   NewPaintSQLite(db),
		Shirt:   NewShirtSQLite(db),
		Order:   NewOrderSQLite(db),
	}
}
