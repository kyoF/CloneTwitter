package infrastructure

import (
	"gorm.io/gorm"

	"backend/domain/entity"
	"backend/domain/repository"
)

type userInfra struct{
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
    return &userInfra{db}
}

func (i *userInfra) FetchUser(userId string) (*entity.User, error) {
    var user *entity.User
    err := i.db.First(&user).Where("userId = ?", userId).Error
    if err != nil {
        return nil, err
    }
    return user, nil
}

