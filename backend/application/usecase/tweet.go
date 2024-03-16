package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type ITweetUsecase interface {
	Create(userId string, text string) (*entity.Tweet, error)
    FetchAll() ([]entity.Tweet, error)
    Fetch(tweetId string) (entity.Tweet, error)
}

type tweetUsecase struct {
	tweetRepo repository.ITweetRepository
}

func NewTweetUsecase(tweetRepo repository.ITweetRepository) ITweetUsecase {
	return &tweetUsecase{
		tweetRepo,
	}
}

func (uc *tweetUsecase) Create(userId string, text string) (*entity.Tweet, error) {
	tweet, err := uc.tweetRepo.Create(userId, text)
	return tweet, err
}

func (uc *tweetUsecase) FetchAll() ([]entity.Tweet, error) {
    tweets, err := uc.tweetRepo.FetchAll()
    return tweets, err
}

func (uc *tweetUsecase) Fetch(tweetId string) (entity.Tweet, error) {
    tweet, err := uc.tweetRepo.Fetch(tweetId)
    return tweet, err
}
