package daoimpl

import (
	"context"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	"log"
	"testing"
	"time"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	return ctx
}

func TestUserLocalImpl_AddUserInfo(t *testing.T) {
	ctx := initenv()
	data:= map[string]interface{}{
		"user_name" : "110",
		"user_description" : " ",
		"e_mail" : "1109942647@qq.com",
		"user_password" : "123456",
		"phone" : " ",
		"token" : " ",
		"created_at" : time.Now().Unix(),
		"created_by" : "zhangwei",
		"updated_at" : time.Now().Unix(),
		"updated_by" : "zhangwei",
		"is_deleted" : false,
	}

	daoimpl.UserDao.AddUserInfo(ctx ,data)
}

func TestUserLocalImpl_UpdateUserInfo(t *testing.T) {
	ctx := initenv()

	cond := map[string]interface{}{
		"id" : 1,
	}
	data := map[string]interface{}{
		"user_name" : "snow",
		"phone": "1555555555",
		"e_mail" : "888@qq.com",
		"description" : "好喜欢你呀",
	}
	query := buildUpdate(data,cond,domain.EmptyUser.TableName())
	log.Printf("%v %s",ctx,query)


}

func TestUserLocalImpl_DelUserInfo(t *testing.T) {

}

func TestUserLocalImpl_ListUserInfo(t *testing.T) {

}