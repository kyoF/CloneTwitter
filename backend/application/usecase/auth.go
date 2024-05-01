package usecase

import (
	"backend/domain/repository"
	"backend/domain/value_objects"
	"log"
)

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
	_, err := u.authRepo.GetByEmail(email)
	if err != nil {
		log.Fatal(err)
	}
	hashedPass, err := value_objects.NewPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	err = u.authRepo.SignUp(userId, email, hashedPass.Value)
	return err
}

func (u *authUsecase) Login(userId string, email string, password string) error {
	err := u.authRepo.Login(userId, email, password)
	return err
}
