package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type ICreateTweetUsecase interface {
    CreateTweet(userId string, text string) (*entity.Tweet, error)
}

type createTweetUsecase struct{
    tweetRepo repository.ITweetRepository
}

func NewCreateTweetUsecase(tweetRepo repository.ITweetRepository) ICreateTweetUsecase {
    return &createTweetUsecase{
        tweetRepo,
    }
}

func (uc *createTweetUsecase) CreateTweet(userId string, text string) (*entity.Tweet, error) {
    tweet, err := uc.tweetRepo.CreateTweet(userId, text)
    if err != nil {
        return nil, err
    }
    return tweet, nil
}

