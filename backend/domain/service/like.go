package service

import "backend/domain/repository"

type ILikeService interface {
    IsExist(userId string, tweetId string) bool
}

type likeService struct {
     likeRepo repository.ILikeRepository
}

func NewLikeService(likeRepo repository.ILikeRepository) ILikeService {
    return &likeService{likeRepo}
}

func (s *likeService) IsExist(userId string, tweetId string) bool {
    like := s.likeRepo.Find(userId, tweetId)
    if like != nil {
        return true
    } else {
        return false
    }
}
