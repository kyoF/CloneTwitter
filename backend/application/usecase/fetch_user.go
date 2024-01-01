package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type IFetchUserUsecase interface {
    FetchUser(userId string) (*entity.User, []entity.Tweet, error)
}

type fetchUserUsecase struct{
    userRepo repository.IUserRepository
    tweetRepo repository.ITweetRepository
}

func NewFetchUserUsecase(userRepo repository.IUserRepository, tweetRepo repository.ITweetRepository) IFetchUserUsecase {
    return &fetchUserUsecase{
        userRepo,
        tweetRepo,
    }
}

func (uc *fetchUserUsecase) FetchUser(userId string) (*entity.User, []entity.Tweet, error) {
    user, err := uc.userRepo.FetchUser(userId)
    if err != nil {
        return nil, nil, err
    }
    tweet, err := uc.tweetRepo.FetchAllUserTweets(user.UserID)
    if err != nil {
        return nil, nil, err
    }
    return user, tweet, nil
}

