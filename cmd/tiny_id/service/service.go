package service

import "context"

type TinyIDService struct {
	ctx context.Context
}

func NewTinyIDService(ctx context.Context) *TinyIDService {
	return &TinyIDService{ctx: ctx}
}
