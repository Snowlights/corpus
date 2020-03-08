package daoimpl

import (
	"context"
	"github.com/Snowlights/corpus/model/dao"
	"log"
)

var (
	UserDao dao.UserDao
)

func PrePare(ctx context.Context){
	fun := "GrpcController.Prepare -->"

	UserDao = NewUserDao()


	log.Printf("%v %s success ",ctx,fun)
	return
}