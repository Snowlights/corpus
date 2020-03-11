package controller

import (
	"context"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"sync"
	"time"
)

type cookie struct {
	mu sync.Mutex
	cookieList []string
}

var cookieList cookie

func AddCookieToList(cookie string) bool{
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()

	cookieList.cookieList = append(cookieList.cookieList,cookie)
	return true
}

func DelCookieFromList(cookie string) bool {
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()
	for i,item := range cookieList.cookieList{
		if item == cookie{
			cookieList.cookieList = append(cookieList.cookieList[:i],cookieList.cookieList[i+1:]...)
		}
	}

	return false
}

func CheckOnLine(cookie string) bool{
	cookieList.mu.Lock()
	defer cookieList.mu.Unlock()

	for _, item := range cookieList.cookieList{
		if item == cookie{
			return true
		}
	}
	return false
}

func AddAuth(ctx context.Context,req *corpus.AddAuthReq) *corpus.AddAuthRes{
	fun := "AddAuth -->"
	res := &corpus.AddAuthRes{}

	data ,conds := toAddAuth(ctx,req)

	dataList ,err := daoimpl.AuthDao.GetAuth(ctx,conds)
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
			Msg:                  "已存在的权限码",
		}
		return res
	}

	lastInsertId , err := daoimpl.AuthDao.AddAuth(ctx,data)
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

func toAddAuth(ctx context.Context,req *corpus.AddAuthReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"auth_code" : req.AuthCode,
		"auth_description" : req.AuthDescription,
		"service_name" : req.ServiceName,
		"created_at" : now,
		"created_by" : req.Cookie,
		"updated_at" : now,
		"updated_by" : req.Cookie,
		"is_deleted" : false,
	}
	conds := map[string]interface{}{
		"auth_code" : req.AuthCode,
		"is_deleted" : false,
	}
	return data,conds
}

func UpdateAuth(ctx context.Context,req *corpus.UpdateAuthReq) *corpus.UpdateAuthRes{
	fun := "UpdateAuth -->"
	res := &corpus.UpdateAuthRes{}

	data, conds := toUpdateAuth(ctx,req)

	rowsAffected, err := daoimpl.AuthDao.UpdateAuth(ctx,data,conds)
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

func toUpdateAuth(ctx context.Context,req *corpus.UpdateAuthReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"auth_code" : req.AuthCode,
		"auth_description" : req.AuthDescription,
		"service_name" : req.ServiceName,
		"updated_at" : now,
		"updated_by" : req.Cookie,
	}
	conds := map[string]interface{}{
		"id" : req.Id,
	}
	return data,conds
}


func DelAuth(ctx context.Context,req* corpus.DelAuthReq) *corpus.DelAuthRes{
	fun := "DelAuth -->"
	res := &corpus.DelAuthRes{}

	data , conds := toDelAuth(ctx,req)

	rowsAffected, err := daoimpl.AuthDao.DelAuth(ctx,data,conds)
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

func toDelAuth(ctx context.Context,req* corpus.DelAuthReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_deleted": true,
		"updated_at" : now,
		"updated_by" : req.Cookie,
	}
	conds := map[string]interface{}{}

	if req.Id != 0{
		conds["id"] = req.Id
	}
	if req.AuthCode != ""{
		conds["auth_code"] = req.AuthCode
	}

	return data,conds
}


func ListAuth(ctx context.Context,req* corpus.ListAuthReq) *corpus.ListAuthRes{
	fun := "ListAuth -->"
	res := &corpus.ListAuthRes{}

	limit, conds := toListAuth(ctx,req)
	dataList, err := daoimpl.AuthDao.ListAuth(ctx,limit,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	total, err := daoimpl.AuthDao.CountAuth(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	data := formCorpusAuthInfoList(dataList)
	res.Data = &corpus.ListAuthData{
		Items:                data,
		Total:                total,
		Offset:               req.Offset + int64(len(dataList)),
		More:                 (req.Offset + int64(len(dataList))) < total,
	}

	return res
}

func formCorpusAuthInfoList(input []*domain.AuthInfo) []*corpus.AuthInfo{
	var output []*corpus.AuthInfo

	for _,item := range input{
		output = append(output,&corpus.AuthInfo{
			Id:                   item.Id,
			AuthName:             item.AuthCode,
			AuthDescription:      item.AuthDescription,
			ServiceName:          item.ServiceName,
		})
	}
	return output
}

func toListAuth(ctx context.Context,req* corpus.ListAuthReq) (map[string]interface{},map[string]interface{}){
	limit := map[string]interface{}{
		"limit" : req.Limit,
		"offset" : req.Offset,
	}
	conds := map[string]interface{}{
		"is_deleted" : false,
	}
	return limit,conds
}

func AddUserAuth(ctx context.Context,req* corpus.AddUserAuthReq) *corpus.AddUserAuthRes{
	res:= &corpus.AddUserAuthRes{}
	return res
}

func DelUserAuth(ctx context.Context,req *corpus.DelUserAuthReq) *corpus.DelUserAuthRes{
	res := &corpus.DelUserAuthRes{}
	return res
}

func ListUserAuth(ctx context.Context, req *corpus.ListUserAuthReq) *corpus.ListUserAuthRes{
	res:= &corpus.ListUserAuthRes{}
	return res
}
