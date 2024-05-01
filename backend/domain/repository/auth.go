package repository

import "backend/domain/entity"

type IAuthRepository interface {
	SignUp(userId, email, password string) error
	Login(userId, email, password string) error
	GetByEmail(email string) (entity.Auth, error)
}
