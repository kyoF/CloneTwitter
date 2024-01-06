package entity

import "github.com/google/uuid"

type Tweet struct {
    TweetID uuid.UUID `json:"tweetId" gorm:"column:tweet_id"`
    UserID string `json:"userId" gorm:"column:user_id"`
    Text string `json:"text" gorm:"column:text"`
    LikesCount int `json:"likesCount" gorm:"column:likes_count"`
}

func NewTweet(userId string, text string) (*Tweet, error) {
    tweetId, err := uuid.NewUUID()
    if err != nil {
        return nil, err
    }
    tweet := &Tweet{
        TweetID: tweetId,
        UserID: userId,
        Text: text,
        LikesCount: 0,
    }
    return tweet, nil
}

