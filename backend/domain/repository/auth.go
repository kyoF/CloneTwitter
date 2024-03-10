package repository

type IAuthRepository interface {
	SignUp(userId string, email string, Password string) error
	Login(userId string, email string, Password string) error
}
