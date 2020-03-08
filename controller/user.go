package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func LoginUser(ctx context.Context,req *corpus.LoginUserReq) *corpus.LoginUserRes{
	res := &corpus.LoginUserRes{}
	return res
}