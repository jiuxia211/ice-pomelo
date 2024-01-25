package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TinyID struct {
	ID        int64
	BizType   int64
	MaxID     int64
	Step      int64
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateTinyID(ctx context.Context, tinyID *TinyID) (err error) {
	return DB.WithContext(ctx).Create(tinyID).Error
}

func CheckTinyIDExist(ctx context.Context, bizType int64) (isExist bool, err error) {
	var tinyID = new(TinyID)
	err = DB.WithContext(ctx).Where("biz_type = ?", bizType).First(&tinyID).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func GetMaxTinyID(ctx context.Context, bizType int64) (maxID int64, err error) {
	var tinyID = new(TinyID)
	err = DB.WithContext(ctx).Where("biz_type = ?", bizType).First(&tinyID).Error
	if err != nil {
		return 0, err
	}
	maxID = tinyID.MaxID
	// 更新tinyID
	tinyID.Version++
	tinyID.MaxID += tinyID.Step
	err = DB.Save(&tinyID).Error
	if err != nil {
		return 0, err
	}

	return maxID, nil
}
