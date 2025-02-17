package models

import (
	"context"
	"siki/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func Create(ctx context.Context, data User) error {
	return configs.Connection.WithContext(ctx).
		Create(&data).Error
	// INSERT INTO users value
}

func Read(ctx context.Context, id uint) (User, error) {
	var user User

	err := configs.Connection.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error

	return user, err
}
func Update(ctx context.Context, id string, data User) error {
	result := configs.Connection.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":  data.Name,
			"email": data.Email,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func Delete(ctx context.Context, id string) error {
	result := configs.Connection.WithContext(ctx).
		Delete(&User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
