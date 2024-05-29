package di

import (
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func User(db *gorm.DB) router.IUserRouter {
	userRepo := mysql.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(userRepo)
	return router.NewUserRouter(usecase)
}
