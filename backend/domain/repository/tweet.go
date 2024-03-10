package repository

import "backend/domain/entity"

type ITweetRepository interface {
    FetchAll() ([]entity.Tweet, error)
    Fetch(tweetId string) (entity.Tweet, error)
    Create(userId string, text string) (*entity.Tweet, error)
}

