package entity

type User struct {
    UserId string `json:"userId" gorm:"column:user_id"`
    Name string `json:"name" gorm:"column:name"`
    Mail string `json:"mail" gorm:"column:mail"`
    Pass string `json:"pass" gorm:"column:pass"`
}

func NewUser(userId string, name string, mail string, pass string) (*User, error) {
    user := &User{
        UserId: userId,
        Name: name,
        Mail: mail,
        Pass: pass,
    }
    return user, nil
}

