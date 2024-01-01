package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/utils"
)

func NewDB() *gorm.DB {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?loc=Asia&&2FTokyo&parseTime=True",
        utils.GetEnv("DB_USER", ""),
        utils.GetEnv("DB_PASSWORD", ""),
        utils.GetEnv("DB_HOST", ""),
        utils.GetEnv("DB_PORT", ""),
        utils.GetEnv("DB_DBNAME", ""),
    )
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    return db
}

