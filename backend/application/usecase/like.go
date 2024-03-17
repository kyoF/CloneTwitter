package usecase

import (
	"backend/domain/repository"
	"backend/domain/service"
	"fmt"
)

type ILikeUsecase interface {
	ToggleLike(userId string, tweetId string) (int, error)
}

type likeUsecase struct {
	likeRepo repository.ILikeRepository
    tweetRepo repository.ITweetRepository
}

func NewLikeUsecase(likeRepo repository.ILikeRepository, tweetRepo repository.ITweetRepository) ILikeUsecase {
	return &likeUsecase{likeRepo, tweetRepo}
}

func (uc *likeUsecase) ToggleLike(userId string, tweetId string) (int, error) {
    var err error

    fmt.Println("aaa")
    likeService := service.NewLikeService(uc.likeRepo)
    if likeService.IsExist(userId, tweetId) {
        err = uc.likeRepo.Delete(userId, tweetId)
    } else {
        err = uc.likeRepo.Create(userId, tweetId)
    }
    fmt.Println("bbb")
    if err != nil {
        return -1, err
    }
    fmt.Println("ccc")

    likes, err := uc.likeRepo.FetchAllByTweetId(tweetId)
    fmt.Println(likes)
    if err != nil {
        return -1, err
    }
    fmt.Println("ddd")
    likesCount := len(likes)

    tweet, err := uc.tweetRepo.Fetch(tweetId)
    if err !=nil {
        return -1, err
    }
    fmt.Println("eee")

    tweet.UpdateCount(likesCount)
    err = uc.tweetRepo.Update(tweet)

    return likesCount, err
}
