package services

import (
	"gin_crud/models"
	"gin_crud/repositories"
)

type UserService interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(id uint, input *models.User) (*models.User, error)
	Delete(id uint) error
}

type service struct {
	repo repositories.UserRepository
}

func NewService(repo repositories.UserRepository) UserService {
	return &service{repo}
}

func (s *service) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *service) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id uint, input *models.User) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Age = input.Age

	err = s.repo.Update(user)
	return user, err
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
