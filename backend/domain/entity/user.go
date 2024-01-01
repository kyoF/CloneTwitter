package entity

import "time"

type User struct {
    ID int
    UserID string
    Name string
    Mail string
    Pass string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewUser(userId string, name string, mail string, pass string) (*User, error) {
    var now time.Time = time.Now()
    user := &User{
        UserID: userId,
        Name: name,
        Mail: mail,
        Pass: pass,
        CreatedAt: now,
    }
    return user, nil
}

