package entity

type Auth struct {
	UserId   string `json:"userId" gorm:"column:user_id"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
