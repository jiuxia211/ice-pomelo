package service

import (
	"github.com/jiuxia211/ice-pomelo/cmd/tiny_id/dal/db"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

func (s *TinyIDService) CreateTinyID(bizType int64) (err error) {

	isExist, err := db.CheckTinyIDExist(s.ctx, bizType)

	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	tinyID := db.TinyID{
		BizType: bizType,
		MaxID:   constants.StartID,
		Step:    constants.IDStep,
		Version: 1,
	}

	err = db.CreateTinyID(s.ctx, &tinyID)
	if err != nil {
		return err
	}

	return nil
}
