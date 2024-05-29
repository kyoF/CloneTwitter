package usecase

import (
	"backend/domain/repository"
	"backend/domain/service"
)

type ILikeUsecase interface {
	ToggleLike(userId, tweetId string) (int, error)
}

type likeUsecase struct {
	likeRepo  repository.ILikeRepository
	tweetRepo repository.ITweetRepository
}

func NewLikeUsecase(likeRepo repository.ILikeRepository, tweetRepo repository.ITweetRepository) ILikeUsecase {
	return &likeUsecase{likeRepo, tweetRepo}
}

func (uc *likeUsecase) ToggleLike(userId, tweetId string) (int, error) {
	var err error

	likeService := service.NewLikeService(uc.likeRepo)
	if likeService.IsExist(userId, tweetId) {
		err = uc.likeRepo.Delete(userId, tweetId)
	} else {
		err = uc.likeRepo.Create(userId, tweetId)
	}
	if err != nil {
		return -1, err
	}

	likes, err := uc.likeRepo.FetchAllByTweetId(tweetId)
	if err != nil {
		return -1, err
	}
	likesCount := len(likes)

	tweet, err := uc.tweetRepo.Fetch(tweetId)
	if err != nil {
		return -1, err
	}

	tweet.UpdateCount(likesCount)
	err = uc.tweetRepo.Update(tweet)

	return likesCount, err
}
