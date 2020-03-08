package daoimpl

import (
	"context"
	"github.com/Snowlights/corpus/model"
	"testing"
	"time"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	return ctx
}

func TestUserLocalImpl_AddUserInfo(t *testing.T) {
	ctx := initenv()
	UserDao := NewUserDao()
	data:= map[string]interface{}{
		"user_name" : "110",
		"e_mail" : "1109942647@qq.com",
		"created_at" : time.Now().Unix(),
		"created_by" : "zhangwei",
		"updated_at" : time.Now().Unix(),
		"updated_by" : "zhangwei",
		"is_deleted" : false,
	}
	UserDao.AddUserInfo(ctx ,data)
}
