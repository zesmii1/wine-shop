package service

import "wine-shop/internal/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	DeleteUserByID(id string) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) DeleteUserByID(id string) error {
	return s.repo.DeleteUserByID(id) // исправлено на s.repo
}
