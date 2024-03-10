package infrastructure

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type authInfra struct {
    db *gorm.DB
}

func NewAuthInfra(db *gorm.DB, authRepo repository.IAuthRepository) repository.IAuthRepository {
	return &authInfra{db: db}
}

func (i *authInfra) Login(userId string, email string, password string) error {
	var user *entity.Auth
	err := i.db.Where("'email' = ? and 'password' = ?", email, password).First(&user).Error
	return err
}

func (i *authInfra) SignUp(userId string, email string, password string) error {
	user := entity.NewAuth(userId, email, password)
	f := func(tx *gorm.DB) error {
		return tx.Create(&user).Error
	}
    err := i.db.Transaction(f)
	return err
}
