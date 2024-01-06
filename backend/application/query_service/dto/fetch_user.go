package dto

import "backend/domain/entity"

type FetchUserDto struct{
    UserId string `json:"userId"`
    Name string `json:"name"`
    Tweets []entity.Tweet `json:"tweets"`
}

func NewFetchUserDto(userId string, name string, tweets []entity.Tweet) *FetchUserDto {
    return &FetchUserDto{
        UserId: userId,
        Name: name,
        Tweets: tweets,
    }
}

