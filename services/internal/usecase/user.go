package usecase

import (
	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/repository"
)

type User interface {
	SignIn(c *domain.Credentials) (*domain.JWT, error)
}

type UserUseCase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) SignIn(c *domain.Credentials) (*domain.JWT, error) {
	accessToken, err := uc.userRepo.SignIn(c)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
