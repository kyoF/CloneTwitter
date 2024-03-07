package usecase

import "backend/domain/repository"

type ISignUpUsecase interface {
	SignUp(userId string, name string, email string, password string) error
}

type signUpUsecase struct {
	userRepo repository.IUserRepository
}

func NewSignUpUsecase(userRepo repository.IUserRepository) ISignUpUsecase {
	return &signUpUsecase{userRepo}
}

func (uc *signUpUsecase) SignUp(userId string, name string, email string, password string) error {
	err := uc.userRepo.SignUp(userId, name, email, password)
	return err
}
