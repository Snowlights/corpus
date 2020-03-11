package controller

import (
	"context"
	"github.com/Snowlights/corpus/model/daoimpl"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"time"
)

func AddAdminUser(ctx context.Context,req* corpus.AddAdminUserReq) *corpus.AddAdminUserRes{
	fun := "AddAdminUser -->"
	res := &corpus.AddAdminUserRes{}
	data, conds := toAddAdminUser(ctx,req)

	dataList ,err := daoimpl.AdminUserDao.GetAdminUser(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	if len(dataList) > 0{
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "已存在的管理用户",
		}
		return res
	}

	lastInsertId ,err := daoimpl.AdminUserDao.AddAdminUser(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func DelAdminUser(ctx context.Context,req*corpus.DelAdminUserReq) *corpus.DelAdminUserRes{
	fun := "DelAdminUser -->"
	res := &corpus.DelAdminUserRes{}

	data, conds := toDelAdminUser(ctx,req)
	rowsAffected , err := daoimpl.AdminUserDao.DelAdminUser(ctx,data,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	log.Printf("%v %v success ,rowsAffected %d",ctx,fun,rowsAffected)
	return res
}

func ListAdminUser(ctx context.Context,req *corpus.ListAdminUserReq) *corpus.ListAdminUserRes{
	fun := "ListAdminUser --> "
	res := &corpus.ListAdminUserRes{}
	limit, conds := toListAdminUser(ctx,req)

	dataList, err := daoimpl.AdminUserDao.ListAdminUser(ctx,limit,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	if len(dataList) == 0{
		res.Data = &corpus.ListAdminUserData{
			Items:                nil,
			Total:                0,
			Offset:               0,
			More:                 false,
		}
		return res
	}

	var idList []int64
	for _, item := range dataList{
		idList = append(idList,item.UserId)
	}
	userConds := map[string]interface{}{
		"id" : idList,
	}
	userDataList, err := daoimpl.UserDao.ListUserInfo(ctx,limit,userConds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	total,err := daoimpl.AdminUserDao.CountAdminUser(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	data := formCorpusUserInfoList(userDataList)

	res.Data = &corpus.ListAdminUserData{
		Items:                data,
		Total:                total,
		Offset:               req.Offset + int64(len(dataList)),
		More:                 (req.Offset + int64(len(dataList))) < total,
	}

	return res
}

func toAddAdminUser(ctx context.Context,req *corpus.AddAdminUserReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"user_id" : req.UserId,
		"created_at" : now,
		"created_by" : req.Cookie,
		"updated_at" : now,
		"updated_by" : req.Cookie,
		"is_deleted" : false,
	}
	conds := map[string]interface{}{
		"user_id" : req.UserId,
		"is_deleted" : false,
	}

	return data,conds
}

func toDelAdminUser(ctx context.Context,req *corpus.DelAdminUserReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_deleted" : true,
		"updated_at" : now,
		"updated_by" : req.Cookie,
	}
	conds := map[string]interface{}{
		"user_id" : req.UserId,
	}
	return data,conds
}

func toListAdminUser(ctx context.Context, req *corpus.ListAdminUserReq) (map[string]interface{}, map[string]interface{}){

	limit := map[string]interface{}{
		"limit" : req.Limit,
		"offset" : req.Offset,
	}

	conds := map[string]interface{}{
		"is_deleted" : false,
	}

	return limit,conds
}