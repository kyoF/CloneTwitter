package repository

import "backend/domain/entity"

type IUserRepository interface {
	Fetch(userId string) (entity.User, error)
}
