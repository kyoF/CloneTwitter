package repository

import "backend/domain/entity"

type ITweetRepository interface {
    FetchAllUserTweets(userId string) ([]entity.Tweet, error)
    FetchAllTweets() ([]entity.Tweet, error)
    CreateTweet(userId string, text string) (*entity.Tweet, error)
}

