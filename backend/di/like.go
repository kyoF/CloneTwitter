package di

import (
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Like(db *gorm.DB) router.ILikeRouter {
	likeRepo := mysql.NewLikeInfra(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepo)
	return router.NewLikeRouter(likeUsecase)
}
