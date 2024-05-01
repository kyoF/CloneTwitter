package value_objects

import "golang.org/x/crypto/bcrypt"

type Password struct {
	Value string
}

func NewPassword(pass string) (Password, error) {
	err := validation(pass)
	if err != nil {
		return Password{}, err
	}
	password := Password{Value: pass}
	return password, nil
}

func validation(_ string) error {
	return nil
}

func (p *Password) Hashed(pass Password) (Password, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass.Value), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	return Password{Value: string(hashedPass)}, nil
}

func (p *Password) Check(existPass, inputPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(existPass), []byte(inputPass))
}
