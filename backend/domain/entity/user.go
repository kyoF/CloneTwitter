package entity

type User struct {
	UserId  string `json:"userId" gorm:"column:user_id"`
	Name    string `json:"name" gorm:"column:name"`
	Profile string `json:"profile" gorm:"column:profile"`
}
