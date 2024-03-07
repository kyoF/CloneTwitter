package entity

type User struct {
    UserId string `json:"userId" gorm:"column:user_id"`
    Name string `json:"name" gorm:"column:name"`
    Mail string `json:"email" gorm:"column:email"`
    Pass string `json:"password" gorm:"column:password"`
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

