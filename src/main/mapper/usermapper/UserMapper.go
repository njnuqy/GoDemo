package usermapper

import (
	"GoDemo/src/main/entity"
	"GoDemo/src/main/mapper"
	"errors"
	"gorm.io/gorm"
)

func FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := mapper.DB.Where("name = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，返回错误或 nil，取决于你的业务逻辑
			return nil, gorm.ErrRecordNotFound
		}
		// 如果发生其他错误，返回该错误
		return nil, result.Error
	}
	return &user, nil
}
