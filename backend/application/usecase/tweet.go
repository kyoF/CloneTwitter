package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type ITweetUsecase interface {
	Create(userId string, text string) (*entity.Tweet, error)
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
