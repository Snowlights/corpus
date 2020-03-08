package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)
type UserDao interface {
	GetUserInfo(ctx context.Context,conds map[string]interface{}) ([]*domain.UserInfo,error)
	AddUserInfo(ctx context.Context,dataList map[string]interface{}) (int64, error)
	UpdateUserInfo(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	DelUserInfo(ctx context.Context, data,conds map[string]interface{})(int64,error)
	ListUserInfo(ctx context.Context,conds map[string]interface{}) ([]*domain.UserInfo,error)
}
