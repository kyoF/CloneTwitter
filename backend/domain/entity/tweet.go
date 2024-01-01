package entity

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
    ID int
    TweetID uuid.UUID
    UserID string
    Text string
    LikesCount int
    CreatedAt time.Time
}

func NewTweet(userId string, text string) (*Tweet, error) {
    tweetId, err := uuid.NewUUID()
    if err != nil {
        return nil, err
    }
    var now time.Time = time.Now()
    tweet := &Tweet{
        TweetID: tweetId,
        UserID: userId,
        Text: text,
        LikesCount: 0,
        CreatedAt: now,
    }
    return tweet, nil
}

