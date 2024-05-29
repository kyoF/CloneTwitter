package repository

import "backend/domain/entity"

type ITweetRepository interface {
	FetchAll() ([]entity.Tweet, error)
	FetchAllByUserId(userId string) ([]entity.Tweet, error)
	Fetch(tweetId string) (entity.Tweet, error)
	Create(userId, text string) (*entity.Tweet, error)
	Update(tweet entity.Tweet) error
}
