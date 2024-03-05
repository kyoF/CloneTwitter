package query_service

import (
	"backend/domain/repository"
)

type ILoginQueryService interface {
    Login(email string, password string) error
}

type loginQueryService struct {
    userRepo repository.IUserRepository
}

func NewLoginQueryService(userRepo repository.IUserRepository) ILoginQueryService {
    return &loginQueryService{
        userRepo,
    }
}

func (qs *loginQueryService) Login(email string, password string) error {
    return qs.userRepo.Login(email, password)
}

