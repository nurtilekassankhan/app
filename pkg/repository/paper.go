package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Paper interface {
	CreatePaper(paper app.Paper) (int, error)
	GetPaperById(id int) (app.Paper, error)
	GetAllPapers() ([]app.Paper, error)
	UpdatePaper(id int, paper app.Paper) error
	DeletePaper(id int) error
}

type PaperSQLite struct {
	db *sqlx.DB
}

func NewPaperSQLite(db *sqlx.DB) *PaperSQLite {
	return &PaperSQLite{db: db}
}

func (r *PaperSQLite) CreatePaper(paper app.Paper) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (type, size) values ($1, $2) returning id", papersTable)
	row := r.db.QueryRow(query, paper.Type, paper.Size)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PaperSQLite) GetPaperById(id int) (app.Paper, error) {
	var paper app.Paper
	query := fmt.Sprintf("select id, type, size FROM %s WHERE id = $1", papersTable)
	if err := r.db.Get(&paper, query, id); err != nil {
		return paper, err
	}
	return paper, nil
}

func (r *PaperSQLite) GetAllPapers() ([]app.Paper, error) {
	var papers []app.Paper
	query := fmt.Sprintf("select id, type, size FROM %s", papersTable)
	if err := r.db.Select(&papers, query); err != nil {
		return nil, err
	}
	return papers, nil
}

func (r *PaperSQLite) UpdatePaper(id int, paper app.Paper) error {
	query := fmt.Sprintf("update %s set type = $1, size = $2 where id = $3", papersTable)
	_, err := r.db.Exec(query, paper.Type, paper.Size, id)
	return err
}

func (r *PaperSQLite) DeletePaper(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", papersTable)
	_, err := r.db.Exec(query, id)
	return err
}
