package service

import (
	"app"
	"app/pkg/repository"
)

type Textile interface {
	CreateTextile(textile app.Textile) (int, error)
	GetAllTextiles() ([]app.Textile, error)
	GetTextileById(id int) (app.Textile, error)
	UpdateTextile(id int, textile app.Textile) error
	DeleteTextile(id int) error
}

type TextileService struct {
	repo repository.Textile
}

func NewTextileService(repo repository.Textile) *TextileService {
	return &TextileService{
		repo: repo,
	}
}

func (s *TextileService) CreateTextile(textile app.Textile) (int, error) {
	return s.repo.CreateTextile(textile)
}

func (s *TextileService) GetAllTextiles() ([]app.Textile, error) {
	return s.repo.GetAllTextiles()
}

func (s *TextileService) GetTextileById(id int) (app.Textile, error) {
	return s.repo.GetTextileById(id)
}

func (s *TextileService) UpdateTextile(id int, textile app.Textile) error {
	return s.repo.UpdateTextile(id, textile)
}

func (s *TextileService) DeleteTextile(id int) error {
	return s.repo.DeleteTextile(id)
}
