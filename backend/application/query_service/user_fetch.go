package query_service

import (
	"backend/application/query_service/dto"
	"backend/domain/repository"
)

type IUserQueryService interface {
    Fetch(userId string) (*dto.UserFetchDto, error)
}

type userQueryService struct{
    userRepo repository.IUserRepository
    tweetRepo repository.ITweetRepository
}

func NewUserQueryService(userRepo repository.IUserRepository, tweetRepo repository.ITweetRepository) IUserQueryService {
    return &userQueryService{
        userRepo,
        tweetRepo,
    }
}

func (qs *userQueryService) Fetch(userId string) (*dto.UserFetchDto, error) {
    user, err := qs.userRepo.Fetch(userId)
    if err != nil {
        return nil, err
    }
    tweets, err := qs.tweetRepo.FetchAllByUserId(userId)
    if err != nil {
        return nil, err
    }
    dto := dto.NewUserFetchDto(
        userId,
        user.Name,
        tweets,
    )
    return dto, nil
}
