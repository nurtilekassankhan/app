package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Textile interface {
	CreateTextile(textile app.Textile) (int, error)
	GetTextileById(id int) (app.Textile, error)
	GetAllTextiles() ([]app.Textile, error)
	UpdateTextile(id int, textile app.Textile) error
	DeleteTextile(id int) error
}

type TextileSQLite struct {
	db *sqlx.DB
}

func NewTextileSQLite(db *sqlx.DB) *TextileSQLite {
	return &TextileSQLite{db: db}
}

func (r *TextileSQLite) CreateTextile(textile app.Textile) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (type) values ($1) returning id", textilesTable)
	row := r.db.QueryRow(query, textile.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TextileSQLite) GetTextileById(id int) (app.Textile, error) {
	var textile app.Textile
	query := fmt.Sprintf("select id, type FROM %s WHERE id = $1", textilesTable)
	if err := r.db.Get(&textile, query, id); err != nil {
		return textile, err
	}
	return textile, nil
}

func (r *TextileSQLite) GetAllTextiles() ([]app.Textile, error) {
	var textiles []app.Textile
	query := fmt.Sprintf("select id, type FROM %s", textilesTable)
	if err := r.db.Select(&textiles, query); err != nil {
		return nil, err
	}
	return textiles, nil
}

func (r *TextileSQLite) UpdateTextile(id int, textile app.Textile) error {
	query := fmt.Sprintf("update %s set type = $1 where id = $2", textilesTable)
	_, err := r.db.Exec(query, textile.Type, id)
	return err
}

func (r *TextileSQLite) DeleteTextile(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", textilesTable)
	_, err := r.db.Exec(query, id)
	return err
}
