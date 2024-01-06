package infrastructure

import (
	"gorm.io/gorm"

	"backend/domain/entity"
	"backend/domain/repository"
)

type tweetInfra struct{
    db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) repository.ITweetRepository {
	return &tweetInfra{db}
}

func (i *tweetInfra) FetchAllUserTweets(userId string) ([]entity.Tweet, error) {
    var tweets []entity.Tweet
    err := i.db.Where("user_id = ?", userId).Find(&tweets).Error
    if err != nil {
        return nil, err
    }
    return tweets, nil
}

func (i *tweetInfra) FetchAllTweets() ([]entity.Tweet, error) {
    var tweets []entity.Tweet
    err := i.db.Find(&tweets).Error
    if err != nil {
        return nil, err
    }
    return tweets, nil
}

func (i *tweetInfra) CreateTweet(userId string, text string) (*entity.Tweet, error) {
    var tweet *entity.Tweet
    err := i.db.Transaction(func(tx *gorm.DB) error {
        tweet, err := entity.NewTweet(userId, text)
        if err != nil {
            return err
        }
        err = i.db.Create(&tweet).Error
        return err
    })
    if err != nil {
        return nil, err
    }
    return tweet, nil
}

