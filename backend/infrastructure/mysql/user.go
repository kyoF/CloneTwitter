package mysql

import (
	"gorm.io/gorm"

	"backend/domain/entity"
	"backend/domain/repository"
)

type userInfra struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userInfra{db}
}

func (i *userInfra) Fetch(userId string) (entity.User, error) {
	var user entity.User
	err := i.db.Model(&entity.User{}).Select("user_id, name, profile").Where("user_id = ?", userId).First(&user).Error
	return user, err
}
