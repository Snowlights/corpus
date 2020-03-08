package controller

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

func LoginUser(ctx context.Context,req *corpus.LoginUserReq) *corpus.LoginUserRes{
	res := &corpus.LoginUserRes{}
	return res
}

func UpdateUserInfo(ctx context.Context ,req *corpus.UpdateUserInfoReq) *corpus.UpdateUserInfoRes{
	res:= &corpus.UpdateUserInfoRes{}
	return res
}

func DelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq) *corpus.DelUserInfoRes{
	res := &corpus.DelUserInfoRes{}
	return res
}

func ListUserInfo(ctx context.Context,req*corpus.ListUserInfoReq) *corpus.ListUserInfoRes{
	res:= &corpus.ListUserInfoRes{}
	return res
}