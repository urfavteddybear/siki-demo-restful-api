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
func Update() {}
func Delete() {}
