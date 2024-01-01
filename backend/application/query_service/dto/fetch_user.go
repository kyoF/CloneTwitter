package dto

import "backend/domain/entity"

type FetchUserDto struct{
    userId string
    name string
    tweets []entity.Tweet
}

func NewFetchUserDto(userId string, name string, tweets []entity.Tweet) *FetchUserDto {
    return &FetchUserDto{
        userId: userId,
        name: name,
        tweets: tweets,
    }
}

