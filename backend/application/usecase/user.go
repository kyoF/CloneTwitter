package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type IUserUsecase interface {
	Fetch(userId string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.IUserRepository
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &userUsecase{userRepo}
}

func (uc *userUsecase) Fetch(userId string) (*entity.User, error) {
	user, err := uc.userRepo.Fetch(userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
