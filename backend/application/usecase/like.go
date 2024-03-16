package usecase

import (
	"backend/domain/repository"
	"backend/domain/service"
)

type ILikeUsecase interface {
	ToggleLike(userId string, tweetId string) error
}

type likeUsecase struct {
	likeRepo repository.ILikeRepository
}

func NewLikeUsecase(likeRepo repository.ILikeRepository) ILikeUsecase {
	return &likeUsecase{likeRepo}
}

func (uc *likeUsecase) ToggleLike(userId string, tweetId string) error {
    likeService := service.NewLikeService(uc.likeRepo)
    if likeService.IsExist(userId, tweetId) {
        return uc.likeRepo.Delete(userId, tweetId)
    } else {
        return uc.likeRepo.Create(userId, tweetId)
    }
}
