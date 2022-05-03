package service

import (
	"app"
	"app/pkg/repository"
)

type Shirt interface {
	CreateShirt(shirt app.Shirt) (int, error)
	GetAllShirts() ([]app.Shirt, error)
	GetShirtById(id int) (app.Shirt, error)
	UpdateShirt(id int, shirt app.Shirt) error
	DeleteShirt(id int) error
}

type ShirtService struct {
	repo repository.Shirt
}

func NewShirtService(repo repository.Shirt) *ShirtService {
	return &ShirtService{
		repo: repo,
	}
}

func (s *ShirtService) CreateShirt(shirt app.Shirt) (int, error) {
	return s.repo.CreateShirt(shirt)
}

func (s *ShirtService) GetAllShirts() ([]app.Shirt, error) {
	return s.repo.GetAllShirts()
}

func (s *ShirtService) GetShirtById(id int) (app.Shirt, error) {
	return s.repo.GetShirtById(id)
}

func (s *ShirtService) UpdateShirt(id int, shirt app.Shirt) error {
	return s.repo.UpdateShirt(id, shirt)
}

func (s *ShirtService) DeleteShirt(id int) error {
	return s.repo.DeleteShirt(id)
}
