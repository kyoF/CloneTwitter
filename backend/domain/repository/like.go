package repository

import "backend/domain/entity"

type ILikeRepository interface {
	Create(userId string, tweetId string) error
	Delete(userId string, tweetId string) error
	Find(userId string, tweetId string) *entity.Like
    FetchesByUserId(userId string) ([]*entity.Like, error)
    FetchAllByTweetId(tweetId string) ([]*entity.Like, error)
}
