package entity

type Auth struct {
	UserId   string `json:"userId" gorm:"column:user_id"`
	Email     string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func NewAuth(userId string, mail string, password string) *Auth {
	auth := &Auth{
		UserId:   userId,
		Email:     mail,
		Password: password,
	}
	return auth
}
