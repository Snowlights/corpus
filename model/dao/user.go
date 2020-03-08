package dao

import "context"

type UserDao interface {
	AddUserInfo(ctx context.Context,dataList map[string]interface{}) (int64, error)
	UpdateUserInfo(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	DelUserInfo(ctx context.Context, data,conds map[string]interface{})(int64,error)
	ListUserInfo(ctx context.Context,conds map[string]interface{}) ([]domain.UserInfo,error)
}
