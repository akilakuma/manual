package service

import (
	"clean-architecture/domain"
	"clean-architecture/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) RegisterUser(email string) error {
	user := domain.NewUser(email)

	return s.Repo.Save(user)
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {

	return s.Repo.FindByEmail(email)
}
