package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type IFetchAllTweetsUsecase interface {
    FetchAllTweets(userId string) ([]entity.Tweet, error)
}

type fetchAllTweetsUsecase struct{
    tweetRepo repository.ITweetRepository
}

func NewFetchAllTweetsUsecase(tweetRepo repository.ITweetRepository) IFetchAllTweetsUsecase {
    return &fetchAllTweetsUsecase{
        tweetRepo,
    }
}

func (uc *fetchAllTweetsUsecase) FetchAllTweets(userId string) ([]entity.Tweet, error) {
    tweets, err := uc.tweetRepo.FetchAllTweets()
    if err != nil {
        return nil, err
    }
    return tweets, nil
}

