package service

import (
	"app"
	"app/pkg/repository"
)

type Textile interface {
	CreateTextile() (int, error)
	GetAllTextile() ([]app.Textile, error)
	GetTextileById(id int) (app.Textile, error)
	UpdateTextile(id int, textile app.Textile) error
	DeletePrinter(id int) error
}

type TextileService struct {
	repo repository.Textile
}

func NewTextileService(repo repository.Textile) *TextileService {
	return &TextileService{
		repo: repo,
	}
}

func (t *TextileService) CreateTextile() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TextileService) GetAllTextile() ([]app.Textile, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TextileService) GetTextileById(id int) (app.Textile, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TextileService) UpdateTextile(id int, textile app.Textile) error {
	//TODO implement me
	panic("implement me")
}

func (t *TextileService) DeletePrinter(id int) error {
	//TODO implement me
	panic("implement me")
}
