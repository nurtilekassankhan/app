package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Printer interface {
	CreatePrinter(printer app.Printer) (int, error)
	GetPrinterById(id int) (app.Printer, error)
	GetAllPrinters() ([]app.Printer, error)
	UpdatePrinter(id int, printer app.Printer) error
	DeletePrinter(id int) error
}

type PrinterSQLite struct {
	db *sqlx.DB
}

func (r *PrinterSQLite) CreatePrinter(printer app.Printer) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (status, efficiency, started_at, expired_at) values ($1, $2, $3, $4) returning id", printersTable)
	row := r.db.QueryRow(query, printer.Status, printer.Efficiency, printer.StartedAt, printer.ExpiredAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PrinterSQLite) GetPrinterById(id int) (app.Printer, error) {
	var printer app.Printer
	query := fmt.Sprintf("select id, status, efficiency, started_at, expired_at FROM %s WHERE id = $1", printersTable)
	if err := r.db.Get(&printer, query, id); err != nil {
		return printer, err
	}
	return printer, nil
}

func (r *PrinterSQLite) GetAllPrinters() ([]app.Printer, error) {
	var printers []app.Printer
	query := fmt.Sprintf("select id, status, efficiency, started_at, expired_at FROM %s", printersTable)
	if err := r.db.Select(&printers, query); err != nil {
		return nil, err
	}
	return printers, nil
}

func (r *PrinterSQLite) UpdatePrinter(id int, printer app.Printer) error {
	query := fmt.Sprintf("update %s set status = $1, efficiency = $2, started_at = $3, expired_at = $4 where id = $5", printersTable)
	_, err := r.db.Exec(query, printer.Status, printer.Efficiency, printer.StartedAt, printer.ExpiredAt, id)
	return err
}

func (r *PrinterSQLite) DeletePrinter(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", printersTable)
	_, err := r.db.Exec(query, id)
	return err
}

func NewPrinterSQLite(db *sqlx.DB) *PrinterSQLite {
	return &PrinterSQLite{db: db}
}
