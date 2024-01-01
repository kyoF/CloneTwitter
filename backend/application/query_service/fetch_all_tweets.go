package query_service

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type IFetchAllTweetsQueryService interface {
    FetchAllTweets(userId string) ([]entity.Tweet, error)
}

type fetchAllTweetsQueryService struct{
    tweetRepo repository.ITweetRepository
}

func NewFetchAllTweetsQueryService(tweetRepo repository.ITweetRepository) IFetchAllTweetsQueryService {
    return &fetchAllTweetsQueryService{
        tweetRepo,
    }
}

func (uc *fetchAllTweetsQueryService) FetchAllTweets(userId string) ([]entity.Tweet, error) {
    tweets, err := uc.tweetRepo.FetchAllTweets()
    if err != nil {
        return nil, err
    }
    return tweets, nil
}

