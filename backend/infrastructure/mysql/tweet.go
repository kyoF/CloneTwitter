package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/domain/entity"
	"backend/domain/repository"
)

type tweetInfra struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) repository.ITweetRepository {
	return &tweetInfra{db}
}

func (i *tweetInfra) FetchAll() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	err := i.db.Find(&tweets).Error
	return tweets, err
}

func (i *tweetInfra) FetchAllByUserId(userId string) ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	err := i.db.Where("user_id = ?", userId).Find(&tweets).Error
	return tweets, err
}

func (i *tweetInfra) Fetch(tweetId string) (entity.Tweet, error) {
	var tweet entity.Tweet
	err := i.db.Where("tweet_id = ?", tweetId).Find(&tweet).Error
	return tweet, err
}

func (i *tweetInfra) Create(userId string, text string) (*entity.Tweet, error) {
	var tweet *entity.Tweet
	tweetId := uuid.NewString()
	err := i.db.Transaction(func(tx *gorm.DB) error {
		tweet = &entity.Tweet{
			TweetId:    tweetId,
			UserId:     userId,
			Text:       text,
			LikesCount: 0,
		}
		err := tx.Create(&tweet).Error
		return err
	})
	return tweet, err
}

func (i *tweetInfra) Update(tweet entity.Tweet) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("tweet_id = ?", tweet.TweetId).Updates(&entity.Tweet{
			LikesCount: tweet.LikesCount,
		}).Error
	})
	return err
}
