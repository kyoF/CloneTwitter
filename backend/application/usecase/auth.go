package usecase

import "backend/domain/repository"

type IAuthUsecase interface {
	SignUp(userId, email, password string) error
	Login(userId, email, password string) error
}

type authUsecase struct {
	authRepo repository.IAuthRepository
}

func NewAuthUsecase(authRepo repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		authRepo,
	}
}

func (u *authUsecase) SignUp(userId, email, password string) error {
	_, err := u.authRepo.GetByEmail(email)
	if err != nil {
		return err
	}
	err = u.authRepo.SignUp(userId, email, password)
	return err
}

func (u *authUsecase) Login(userId, email, password string) error {
	err := u.authRepo.Login(userId, email, password)
	return err
}
