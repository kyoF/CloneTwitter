package usecase

import "backend/domain/repository"

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
    isExist, err := uc.likeRepo.IsExist(userId, tweetId)
    if err != nil {
        return err
    }
    if isExist {
        return uc.likeRepo.Delete(userId, tweetId)
    } else {
        return uc.likeRepo.Create(userId, tweetId)
    }
}
