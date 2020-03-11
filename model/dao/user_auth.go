package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)

type UserAuthDao interface {
	GetUserAuth(ctx context.Context,conds map[string]interface{}) ([]*domain.UserAuthInfo,error)
	AddUserAuth(ctx context.Context,data map[string]interface{}) (int64,error)
	DelUserAuth(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	ListUserAuth(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.UserAuthInfo,error)
	CountUserAuth(ctx context.Context,conds map[string]interface{}) (int64,error)
}