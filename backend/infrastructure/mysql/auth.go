package mysql

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type authInfra struct {
	db *gorm.DB
}

func NewAuthInfra(db *gorm.DB) repository.IAuthRepository {
	return &authInfra{db: db}
}

func (i *authInfra) Login(userId string, email string, password string) error {
	var user *entity.Auth
	err := i.db.Where("'email' = ? and 'password' = ?", email, password).First(&user).Error
	return err
}

func (i *authInfra) SignUp(userId string, email string, password string) error {
	user := entity.Auth{
		UserId:   userId,
		Email:    email,
		Password: password,
	}
	err := i.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&user).Error
	})
	return err
}

func (i *authInfra) GetByEmail(email string) (entity.Auth, error) {
	var auth entity.Auth
	err := i.db.Where("email = ?", email).First(&auth).Error
	return auth, err
}
