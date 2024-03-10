package query_service

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type ITweetQueryService interface {
    FetchAll() ([]entity.Tweet, error)
    Fetch(tweetId string) (entity.Tweet, error)
}

type tweetQueryService struct {
    tweetRepo repository.ITweetRepository
}

func NewTweetQueryService(tweetRepo repository.ITweetRepository) ITweetQueryService {
    return &tweetQueryService{tweetRepo}
}

func (q *tweetQueryService) FetchAll() ([]entity.Tweet, error) {
    tweets, err := q.tweetRepo.FetchAll()
    return tweets, err
}

func (q *tweetQueryService) Fetch(tweetId string) (entity.Tweet, error) {
    tweet, err := q.tweetRepo.Fetch(tweetId)
    return tweet, err
}
