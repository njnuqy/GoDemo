package mapper

import (
	"GoDemo/src/main/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	// 自动迁移 schema，创建或更新 User 和 Item 表
	db.AutoMigrate(entity.User{})
	return nil
}
