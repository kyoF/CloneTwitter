package query_service

import (
	"backend/domain/repository"
    "backend/application/query_service/dto"
)

type IFetchUserQueryService interface {
    FetchUser(userId string) (*dto.FetchUserDto, error)
}

type fetchUserQueryService struct {
    userRepo repository.IUserRepository
    tweetRepo repository.ITweetRepository
}

func NewFetchUserQueryService(userRepo repository.IUserRepository, tweetRepo repository.ITweetRepository) IFetchUserQueryService {
    return &fetchUserQueryService{
        userRepo,
        tweetRepo,
    }
}

func (qs *fetchUserQueryService) FetchUser(userId string) (*dto.FetchUserDto, error) {
    user, err := qs.userRepo.FetchUser(userId)
    if err != nil {
        return nil, err
    }
    tweets, err := qs.tweetRepo.FetchAllUserTweets(user.UserID)
    if err != nil {
        return nil, err
    }
    dto := dto.NewFetchUserDto(
        userId,
        user.Name,
        tweets,
    )
    return dto, nil
}

