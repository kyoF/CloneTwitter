package dto

import "backend/domain/entity"

type UserFetchDto struct {
	UserId string         `json:"userId"`
	Name   string         `json:"name"`
	Tweets []entity.Tweet `json:"tweets"`
}

func NewUserFetchDto(userId string, name string, tweets []entity.Tweet) *UserFetchDto {
	return &UserFetchDto{
		UserId: userId,
		Name:   name,
		Tweets: tweets,
	}
}
