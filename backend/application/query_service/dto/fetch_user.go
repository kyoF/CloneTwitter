package dto

import "backend/domain/entity"

type FetchUserDto struct{
    UserId string
    Name string
    Tweets []entity.Tweet
}

func NewFetchUserDto(userId string, name string, tweets []entity.Tweet) *FetchUserDto {
    return &FetchUserDto{
        UserId: userId,
        Name: name,
        Tweets: tweets,
    }
}

