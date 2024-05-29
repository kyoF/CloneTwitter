package entity

type Tweet struct {
	TweetId    string `json:"tweetId" gorm:"column:tweet_id"`
	UserId     string `json:"userId" gorm:"column:user_id"`
	Text       string `json:"text" gorm:"column:text"`
	LikesCount int    `json:"likesCount" gorm:"column:likes_count"`
}

func (t *Tweet) UpdateCount(likesCount int) {
	t.LikesCount = likesCount
}
