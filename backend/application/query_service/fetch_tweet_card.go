package query_service

import (
	"backend/application/query_service/dto"
	"backend/domain/repository"
)

type IFetchTweetCardQueryService interface {
	FetchAllByUserId(userId string) ([]dto.TweetCard, error)
}

type fetchTweetCardQueryService struct {
	userRepo  repository.IUserRepository
	tweetRepo repository.ITweetRepository
}

func NewFetchTweetCardQueryService(
	userRepo repository.IUserRepository,
	tweetRepo repository.ITweetRepository,
) IFetchTweetCardQueryService {
	return &fetchTweetCardQueryService{
		userRepo,
		tweetRepo,
	}
}

func (qs *fetchTweetCardQueryService) FetchAllByUserId(userId string) ([]dto.TweetCard, error) {
	user, err := qs.userRepo.Fetch(userId)
	if err != nil {
		return nil, err
	}

	tweets, err := qs.tweetRepo.FetchAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	tweetCardDto := make([]dto.TweetCard, 0, len(tweets))
	for _, tweet := range tweets {
		tweetCardDto = append(tweetCardDto, dto.TweetCard{
			UserId:     tweet.UserId,
			Name:       user.UserId,
			TweetId:    tweet.TweetId,
			LikesCount: tweet.LikesCount,
			Text:       tweet.Text,
		})
	}

	return tweetCardDto, nil
}
