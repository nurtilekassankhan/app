package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Paint interface {
	CreatePaint(paint app.Paint) (int, error)
	GetPaintById(id int) (app.Paint, error)
	GetAllPaints() ([]app.Paint, error)
	UpdatePaint(id int, paint app.Paint) error
	DeletePaint(id int) error
}

type PaintSQLite struct {
	db *sqlx.DB
}

func NewPaintSQLite(db *sqlx.DB) *PaintSQLite {
	return &PaintSQLite{db: db}
}

func (r *PaintSQLite) CreatePaint(paint app.Paint) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (type, color, volume) values ($1, $2, $3) returning id", paintsTable)
	row := r.db.QueryRow(query, paint.Type, paint.Color, paint.Volume)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PaintSQLite) GetPaintById(id int) (app.Paint, error) {
	var paint app.Paint
	query := fmt.Sprintf("select id, type, color, volume FROM %s WHERE id = $1", paintsTable)
	if err := r.db.Get(&paint, query, id); err != nil {
		return paint, err
	}
	return paint, nil
}

func (r *PaintSQLite) GetAllPaints() ([]app.Paint, error) {
	var paints []app.Paint
	query := fmt.Sprintf("select id, type, color, volume FROM %s", paintsTable)
	if err := r.db.Select(&paints, query); err != nil {
		return nil, err
	}
	return paints, nil
}

func (r *PaintSQLite) UpdatePaint(id int, paint app.Paint) error {
	query := fmt.Sprintf("update %s set type = $1, color = $2, volume = $3 where id = $4", paintsTable)
	_, err := r.db.Exec(query, paint.Type, paint.Color, paint.Volume, id)
	return err
}

func (r *PaintSQLite) DeletePaint(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", paintsTable)
	_, err := r.db.Exec(query, id)
	return err
}
