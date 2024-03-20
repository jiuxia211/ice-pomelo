package db

import (
	"context"
	"sort"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        int64
	UID       int64
	ToUID     int64
	Content   string
	Read      bool // 已读
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateMessage(message *Message) (err error) {
	err = DB.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUnreadMessage(ctx context.Context, uid int64) (msgList []Message, err error) {
	err = DB.WithContext(ctx).Where("`read` = 0 and to_uid = ?", uid).Find(&msgList).Error
	if err != nil {
		return nil, err
	}
	err = DB.WithContext(ctx).Where("`read` = 0  and to_uid = ?", uid).Update("read", "1").Error
	if err != nil {
		return nil, err
	}
	return msgList, nil
}
func GetAllMessage(ctx context.Context, uid int64, toUid int64) (msgList []Message, err error) {
	msgList1 := make([]Message, 0)
	msgList2 := make([]Message, 0)
	err = DB.WithContext(ctx).Where("uid = ? and to_uid = ?", uid, toUid).Find(&msgList1).Error
	if err != nil {
		return nil, err
	}
	err = DB.WithContext(ctx).Where("uid = ? and to_uid = ?", toUid, uid).Find(&msgList2).Error
	if err != nil {
		return nil, err
	}

	msgList = append(msgList, msgList1...)
	msgList = append(msgList, msgList2...)
	sort.SliceStable(msgList, func(i, j int) bool {
		return msgList[i].CreatedAt.Before(msgList[j].CreatedAt)
	})

	return msgList, nil
}
