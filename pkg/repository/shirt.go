package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Shirt interface {
	CreateShirt(shirt app.Shirt) (int, error)
	GetShirtById(id int) (app.Shirt, error)
	GetAllShirts() ([]app.Shirt, error)
	UpdateShirt(id int, shirt app.Shirt) error
	DeleteShirt(id int) error
}

type ShirtSQLite struct {
	db *sqlx.DB
}

func NewShirtSQLite(db *sqlx.DB) *ShirtSQLite {
	return &ShirtSQLite{db: db}
}

func (r *ShirtSQLite) CreateShirt(shirt app.Shirt) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (size, paint_id, paper_id, textile_id) values ($1, $2, $3, $4) returning id", shirtsTable)
	row := r.db.QueryRow(query, shirt.Size, shirt.PaintId, shirt.PaperId, shirt.TextileId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ShirtSQLite) GetShirtById(id int) (app.Shirt, error) {
	var shirt app.Shirt
	query := fmt.Sprintf("select id, size, paint_id, paper_id, textile_id FROM %s WHERE id = $1", shirtsTable)
	if err := r.db.Get(&shirt, query, id); err != nil {
		return shirt, err
	}
	return shirt, nil
}

func (r *ShirtSQLite) GetAllShirts() ([]app.Shirt, error) {
	var shirts []app.Shirt
	query := fmt.Sprintf("select id, size, paint_id, paper_id, textile_id FROM %s", shirtsTable)
	if err := r.db.Select(&shirts, query); err != nil {
		return nil, err
	}
	return shirts, nil
}

func (r *ShirtSQLite) UpdateShirt(id int, shirt app.Shirt) error {
	query := fmt.Sprintf("update %s set size = $1, paint_id = $2, paper_id = $3, textile_id = $4 where id = $5", shirtsTable)
	_, err := r.db.Exec(query, shirt.Size, shirt.PaintId, shirt.PaperId, shirt.TextileId, id)
	return err
}

func (r *ShirtSQLite) DeleteShirt(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", shirtsTable)
	_, err := r.db.Exec(query, id)
	return err
}