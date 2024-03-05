package repository

import "backend/domain/entity"

type IUserRepository interface {
	FetchUser(userId string) (*entity.User, error)
	SignUp(userId string, name string, email string, password string) error
    Login(email string, password string) error
}
