package service

import "github.com/jiuxia211/ice-pomelo/cmd/tiny_id/dal/db"

func (s *TinyIDService) GetMaxID(bizType int64) (maxID int64, err error) {
	maxID, err = db.GetMaxTinyID(s.ctx, bizType)

	if err != nil {
		return 0, err
	}

	return maxID, err
}
