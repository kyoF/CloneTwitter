package repository

import "backend/domain/entity"

type ILikeRepository interface {
	Create(userId, tweetId string) error
	Delete(userId, tweetId string) error
	Find(userId, tweetId string) *entity.Like
	FetchesByUserId(userId string) ([]*entity.Like, error)
	FetchAllByTweetId(tweetId string) ([]*entity.Like, error)
}
