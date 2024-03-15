package service

import "backend/domain/repository"

type likeService struct {
     likeRepo repository.ILikeRepository
}

func (s *likeService) IsExist(userId string, tweetId string) bool {
    like := s.likeRepo.Find(userId, tweetId)
    if like != nil {
        return true
    } else {
        return false
    }
}
