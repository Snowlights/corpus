package daoimpl

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/domain"
	"log"
	"testing"
	"time"
)

func initenv() context.Context{
	ctx := context.Background()
	model.Prepare(ctx)
	PrePare(ctx)
	return ctx
}

func TestUserLocalImpl_AddUserInfo(t *testing.T) {
	ctx := initenv()
	data:= map[string]interface{}{
		"user_name" : "110",
		"e_mail" : "1109942647@qq.com",
		"user_password" : "123456",
		"created_at" : time.Now().Unix(),
		"created_by" : "zhangwei",
		"updated_at" : time.Now().Unix(),
		"updated_by" : "zhangwei",
		"is_deleted" : false,
	}

	UserDao.AddUserInfo(ctx ,data)
}

func TestUserLocalImpl_UpdateUserInfo(t *testing.T) {
	ctx := initenv()

	cond := map[string]interface{}{
		"id" : 2,
	}
	data := map[string]interface{}{
		"user_name" : "snow",
		"phone": "1555555555",
		"e_mail" : "888@qq.com",
		"user_description" : "好喜欢你呀",
	}
	query := buildUpdate(data,cond,domain.EmptyUser.TableName())
	log.Printf("%v %s",ctx,query)

	UserDao.UpdateUserInfo(ctx,data,cond)
}

func TestUserLocalImpl_DelUserInfo(t *testing.T) {
	ctx := initenv()
	cond := map[string]interface{}{
		"id" : 2,
	}
	data := map[string]interface{}{
		"updated_at" : time.Now().Unix(),
		"updated_by" : "admin",
		"is_deleted" : true,
	}
	query := buildDelete(data,cond,domain.EmptyUser.TableName())
	log.Printf("%v %s",ctx,query)

	UserDao.UpdateUserInfo(ctx,data,cond)

}

func TestUserLocalImpl_ListUserInfo(t *testing.T) {
	ctx := initenv()
	limit := map[string]interface{}{
		"limit" : 5,
		"offset" : 0,
	}
	cond := map[string]interface{}{
		"e_mail" : "888@qq.com",
		//"is_deleted" : false,
	}
	query := buildList(limit,cond,domain.EmptyUser.TableName())

	log.Printf("%v %s",ctx,query)

	r, err := UserDao.ListUserInfo(ctx,limit,cond)
	if err != nil{
		return
	}

	log.Printf("%v",r)
	for _, item := range r{
		fmt.Printf("%v",item)
	}

}