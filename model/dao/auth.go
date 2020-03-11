package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)

type AuthDao interface {
	GetAuth(ctx context.Context,conds map[string]interface{}) ([]*domain.AuthInfo,error)
	AddAuth(ctx context.Context,data map[string]interface{}) (int64,error)
	UpdateAuth(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	DelAuth(ctx context.Context,data, conds map[string]interface{}) (int64,error)
	ListAuth(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.AuthInfo,error)
	CountAuth(ctx context.Context,conds map[string]interface{}) (int64,error)
}