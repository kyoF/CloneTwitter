package entity

import "github.com/google/uuid"

type Tweet struct {
	TweetId    uuid.UUID `json:"tweetId" gorm:"column:tweet_id"`
	UserId     string    `json:"userId" gorm:"column:user_id"`
	Text       string    `json:"text" gorm:"column:text"`
	LikesCount int       `json:"likesCount" gorm:"column:likes_count"`
}

func NewTweet(userId string, text string) (*Tweet, error) {
	tweetId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	tweet := &Tweet{
		TweetId:    tweetId,
		UserId:     userId,
		Text:       text,
		LikesCount: 0,
	}
	return tweet, nil
}
