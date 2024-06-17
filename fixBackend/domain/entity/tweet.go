package entity

type Tweet struct {
	TweetId   string
	UserId    string
	Text      string
	CreatedAt int64
	LikeCount int16
}
