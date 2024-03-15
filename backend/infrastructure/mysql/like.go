package mysql

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"errors"

	"gorm.io/gorm"
)

type likeInfra struct {
	db *gorm.DB
}

func NewLikeInfra(db *gorm.DB) repository.ILikeRepository {
	return &likeInfra{db}
}

func (i *likeInfra) Create(userId string, tweetId string) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		like := &entity.Like{
			UserId:  userId,
			TweetId: tweetId,
		}
		err := i.db.Create(&like).Error
		return err
	})
	return err
}

func (i *likeInfra) Delete(userId string, tweetId string) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		return i.db.Where("user_id = ? AND tweet_id = ?", userId, tweetId).Delete(&entity.Like{}).Error
	})
	return err
}

func (i *likeInfra) Find(userId string, tweetId string) *entity.Like {
	var like *entity.Like
    // gorm returns the state of no record as an ERROR
	_ = i.db.Where("user_id = ? AND tweet_id = ?", userId, tweetId).First(&like)
	return like
}
