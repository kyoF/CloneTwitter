package entity

type User struct {
	UserId  string `json:"userId" gorm:"column:user_id"`
	Name    string `json:"name" gorm:"column:name"`
	Profile string `json:"profile" gorm:"column:profile"`
}

func NewUser(userId string, name string, profile string) (*User, error) {
	user := &User{
		UserId:  userId,
		Name:    name,
		Profile: profile,
	}
	return user, nil
}
