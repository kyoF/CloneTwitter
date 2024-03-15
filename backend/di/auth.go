package di

import (
	"backend/application/usecase"
	"backend/infrastructure"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Auth(db *gorm.DB) router.IAuthRouter {
	authRepo := infrastructure.NewAuthInfra(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	return router.NewAuthRouter(authUsecase)
}
