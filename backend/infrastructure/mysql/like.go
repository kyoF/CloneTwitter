package mysql

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type likeInfra struct {
	db *gorm.DB
}

func NewLikeInfra(db *gorm.DB) repository.ILikeRepository {
	return &likeInfra{db}
}

func (i *likeInfra) Create(userId string, tweetId string) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		like := &entity.Like{
			UserId:  userId,
			TweetId: tweetId,
		}
		err := i.db.Create(&like).Error
		return err
	})
	return err
}

func (i *likeInfra) Delete(userId string, tweetId string) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		return i.db.Where("user_id = ? AND user_id = ?", userId, tweetId).Delete(&entity.Like{}).Error
	})
	return err
}

func (i *likeInfra) IsExist(userId string, tweetId string) (bool, error) {
	var like entity.Like
	result := i.db.Where("user_id = ? AND user_id = ?", userId, tweetId).First(&like)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
