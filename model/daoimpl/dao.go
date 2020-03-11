package daoimpl

import (
	"context"
	"github.com/Snowlights/corpus/model/dao"
	"log"
)

var (
	UserDao dao.UserDao
	AdminUserDao dao.AdminUserDao
	AuthDao dao.AuthDao
	UserAuth dao.UserAuthDao
)

func PrePare(ctx context.Context){
	fun := "GrpcController.Prepare -->"

	UserDao = NewUserDao()
	AdminUserDao = NewAdminUserDao()
	AuthDao = NewAuthrDao()
	UserAuth = NewUserAuthDao()
	log.Printf("%v %s success ",ctx,fun)
	return
}