package dto

type TweetCard struct {
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	TweetId    string `json:"tweetId"`
	LikesCount int    `json:"likesCount"`
	Text       string `json:"text"`
}
