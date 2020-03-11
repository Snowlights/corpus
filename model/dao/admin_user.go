package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)

type AdminUserDao interface {
	GetAdminUser(ctx context.Context,conds map[string]interface{}) ([]*domain.AdminUserInfo,error)
	AddAdminUser(ctx context.Context,data map[string]interface{}) (int64,error)
	DelAdminUser(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	ListAdminUser(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.AdminUserInfo,error)
	CountAdminUser(ctx context.Context,conds map[string]interface{}) (int64,error)
}