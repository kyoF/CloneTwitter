package repository

import "backend/domain/entity"

type ILikeRepository interface {
	Create(userId string, tweetId string) error
	Delete(userId string, tweetId string) error
	Find(userId string, tweetId string) *entity.Like
}
