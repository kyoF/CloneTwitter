package models

type User struct {
	UserId       string
	Name         string
	Introduction string
	Email        string
	Password     string
	CreatedAt    int64
	UpdatedAt    int64
}
