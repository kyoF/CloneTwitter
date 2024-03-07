package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/utils"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEnv("DB_USER", "twitter"),
		utils.GetEnv("DB_PASSWORD", "twitter"),
		utils.GetEnv("DB_HOST", "database"),
		utils.GetEnv("DB_PORT", "3306"),
		utils.GetEnv("DB_NAME", "twitter"),
	)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
