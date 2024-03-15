package repository

type ILikeRepository interface {
	Create(userId string, tweetId string) error
	Delete(userId string, tweetId string) error
	IsExist(userId string, tweetId string) (bool, error)
}
