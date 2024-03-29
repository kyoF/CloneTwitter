package usecase

import "backend/domain/repository"

type IAuthUsecase interface {
	SignUp(userId string, email string, password string) error
	Login(userId string, email string, password string) error
}

type authUsecase struct {
	authRepo repository.IAuthRepository
}

func NewAuthUsecase(authRepo repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		authRepo,
	}
}

func (u *authUsecase) SignUp(userId string, email string, password string) error {
    err := u.authRepo.SignUp(userId, email, password)
    return err
}

func (u *authUsecase) Login(userId string, email string, password string) error {
    err := u.authRepo.Login(userId, email, password)
    return err
}
