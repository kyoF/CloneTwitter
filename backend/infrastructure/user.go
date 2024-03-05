package infrastructure

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

func (i *userInfra) FetchUser(userId string) (*entity.User, error) {
	var user *entity.User
	err := i.db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (i *userInfra) SignUp(userId string, name string, email string, password string) error {
	user, err := entity.NewUser(userId, name, email, password)
	if err != nil {
		return err
	}
	f := func(tx *gorm.DB) error {
        return tx.Create(&user).Error
	}
    return i.db.Transaction(f)
}
