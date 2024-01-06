package entity

type User struct {
    UserID string `json:"userId" gorm:"column:user_id"`
    Name string `json:"name" gorm:"column:name"`
    Mail string `json:"mail" gorm:"column:mail"`
    Pass string `json:"pass" gorm:"column:pass"`
}

func NewUser(userId string, name string, mail string, pass string) (*User, error) {
    user := &User{
        UserID: userId,
        Name: name,
        Mail: mail,
        Pass: pass,
    }
    return user, nil
}

