package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateUser(ctx context.Context, user *User) (userResp *User, err error) {
	err = DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func GetUserByUsername(ctx context.Context, username string) (user *User, err error) {
	user = new(User)

	err = DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(ctx context.Context, uid int64) (user *User, err error) {
	user = new(User)

	err = DB.WithContext(ctx).First(&user, uid).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
