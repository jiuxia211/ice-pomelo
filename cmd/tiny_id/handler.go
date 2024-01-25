package main

import (
	"context"

	"github.com/jiuxia211/ice-pomelo/cmd/tiny_id/pack"
	"github.com/jiuxia211/ice-pomelo/cmd/tiny_id/service"
	tiny_id "github.com/jiuxia211/ice-pomelo/kitex_gen/tiny_id"
)

// TinyIDServiceImpl implements the last service interface defined in the IDL.
type TinyIDServiceImpl struct{}

// GetMaxID implements the TinyIDServiceImpl interface.
func (s *TinyIDServiceImpl) GetMaxID(ctx context.Context, req *tiny_id.GetMaxIDRequest) (resp *tiny_id.GetMaxIDResponse, err error) {
	resp = new(tiny_id.GetMaxIDResponse)

	maxID, err := service.NewTinyIDService(ctx).GetMaxID(req.BizType)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.MaxID = maxID

	return
}
