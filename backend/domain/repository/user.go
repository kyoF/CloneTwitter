package repository

import "backend/domain/entity"

type IUserRepository interface {
    FetchUser(userId string) (*entity.User, error)
}

