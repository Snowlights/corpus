package controller

import (
	"context"
	"fmt"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"time"
)

func LoginUser(ctx context.Context,req *corpus.LoginUserReq) *corpus.LoginUserRes{
	fun := "Controller.LoginUser -- >"
	res := &corpus.LoginUserRes{}

	data,conds,audit := toLoginUser(ctx,req)

	dataList ,err := daoimpl.UserDao.GetUserInfo(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	if len(dataList) > 0 {
		if dataList[0].UserPassword != req.UserPassword{
			res.Errinfo = &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  "密码不正确",
			}
			return res
		}
		cache.AddCookieToList(dataList[0].Token)
		if req.Phone != ""{
			cache.DelPhoneCode(req.Phone)
		}
		return res
	}

	lastInsertId , err := daoimpl.UserDao.AddUserInfo(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	if req.EMail != ""{
		SendEmail(ctx,[]string{req.EMail})
	}

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func LoginOutUserInfo(ctx context.Context, req *corpus.LogoutUserInfoReq) *corpus.LogoutUserInfoRes{
	fun := "Controller.LoginOutUserInfo -- >"
	res := &corpus.LogoutUserInfoRes{}

	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	audit := formLoginOutAudit(req)
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	pass = cache.DelCookieFromList(req.Cookie)
	if !pass{
		log.Fatalf("%v %s 删除失败 %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未找到cookie信息",
		}
		return res
	}

	return res
}

func UpdateUserInfo(ctx context.Context ,req *corpus.UpdateUserInfoReq) *corpus.UpdateUserInfoRes{
	fun := "Controller.UpdateUserInfo -- >"
	res:= &corpus.UpdateUserInfoRes{}

	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	conds, data, audit := toUpdateUserInfo(ctx,req)
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	rowsAffected, err := daoimpl.UserDao.UpdateUserInfo(ctx,data,conds)
	if err != nil{
		log.Fatalf("%v %s error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	log.Printf("%v %v success ,rowsAffected %d",ctx,fun,rowsAffected)
	return res
}

func DelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq) *corpus.DelUserInfoRes{
	fun := "Controller.DelUserInfo -- >"
	res := &corpus.DelUserInfoRes{}

	pass := cache.CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}
	conds, data,audit := toDelUserInfo(ctx,req)
	auditLastInsertId, err := addAudit(ctx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	rowsAffected, err := daoimpl.UserDao.DelUserInfo(ctx,data,conds)
	if err != nil{
		log.Fatalf("%v %s error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	log.Printf("%v %v success ,rowsAffected %d",ctx,fun,rowsAffected)
	return res
}

func ListUserInfo(ctx context.Context,req*corpus.ListUserInfoReq) *corpus.ListUserInfoRes{
	fun := "Controller.ListUserInfo -- >"
	res:= &corpus.ListUserInfoRes{}

	conds, limit := toListUserInfo(ctx,req)

	dataList, err := daoimpl.UserDao.ListUserInfo(ctx,limit,conds)
	if err != nil{
		log.Fatalf("%v %s error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	data := formCorpusUserInfoList(dataList)
	total,err := daoimpl.UserDao.CountUserInfo(ctx,conds)
	if err != nil{
		log.Fatalf("%v %s error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}

	res.Data = &corpus.ListUserInfoData{
		Items:                data,
		Total:                total,
		Offset:               req.Offset + int64(len(data)),
		More:                 total > req.Offset + int64(len(data)),
	}

	return res
}

func formCorpusUserInfoList(input []*domain.UserInfo) []*corpus.UserInfo{

	var output []*corpus.UserInfo

	for _, item := range input{
		output = append(output,&corpus.UserInfo{
			UserId:               item.Id,
			UserName:             item.UserName,
			Phone:                item.Phone,
			EMail:                item.E_mail,
		})
	}
	return output
}

func toLoginUser(ctx context.Context, req *corpus.LoginUserReq) (map[string]interface{}, map[string]interface{}, map[string]interface{}){
	now := time.Now().Unix()
	cookie := RandCookie()
	data := map[string]interface{}{
		"token" : cookie,
		"user_description" : "",
		"created_at" : now,
		"created_by" : "admin",
		"updated_at" : now,
		"updated_by" : "admin",
	}
	conds := map[string]interface{}{
		"is_deleted" : false,
	}
	audit:= map[string]interface{}{
		"table_name" : domain.EmptyUser.TableName(),
		"created_at" : now,
		"is_deleted" : false,
	}
	if req.Phone != ""{
		data["user_name"] = req.Phone
		data["phone"] = req.Phone
		conds["phone"] = req.Phone
		audit["created_by"] = req.Phone
		audit["history"] = fmt.Sprintf("%v user login time %v cookie %v",req.Phone,now,cookie)
		audit["activity"] = fmt.Sprintf("%v user login",req.Phone)
	}

	if req.EMail != ""{
		data["user_name"] = req.EMail
		data["e_mail"] = req.EMail
		data["user_password"] = req.UserPassword
		conds["e_mail"] = req.EMail
		audit["created_by"] = req.EMail
		audit["history"] = fmt.Sprintf("%v user login time %v cookie %v",req.EMail,now,cookie)
		audit["activity"] = fmt.Sprintf("%v user login",req.EMail)
	}
	audit["content"] = fmt.Sprintf("%v", data)

	return data,conds,audit
}

func formLoginOutAudit(req *corpus.LogoutUserInfoReq) map[string]interface{}{
	now := time.Now().Unix()
	audit := map[string]interface{}{
		"table_name" : "user_login_out",
		"content" : fmt.Sprintf("user login out"),
		"history" : fmt.Sprintf("user login out time %v cookie %v",now,req.Cookie),
		"activity" : fmt.Sprintf("%v user login out",req.Cookie),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return audit
}

func toUpdateUserInfo(ctx context.Context,req *corpus.UpdateUserInfoReq) (map[string]interface{},map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	conds := map[string]interface{}{
		"id" : req.UserId,
	}
	data := map[string]interface{}{
		"user_name": req.UserName,
		"phone" : req.Phone,
		"e_mail" : req.EMail,
		"user_description" : req.Description,
		"updated_at" : now,
		"updated_by" : req.UserName,
	}

	audit := map[string]interface{}{
		"table_name" : domain.EmptyUser.TableName(),
		"content" : fmt.Sprintf("%v",data),
		"history" : fmt.Sprintf("user update info time %v cookie %v",now,req.Cookie),
		"activity" : fmt.Sprintf("%v user update info",req.Cookie),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}

	return conds,data,audit
}

func toDelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq) (map[string]interface{},map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	conds := map[string]interface{}{
		"id" : req.UserId,
	}
	data := map[string]interface{}{
		"is_deleted" : true,
		"updated_at" : now,
		"updated_by" : "admin",
	}
	audit := map[string]interface{}{
		"table_name" : domain.EmptyUser.TableName(),
		"content" : fmt.Sprintf("%v",data),
		"history" : fmt.Sprintf("del user info time %v cookie %v",now,req.Cookie),
		"activity" : fmt.Sprintf("%v del user info",req.Cookie),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return conds,data,audit
}

func toListUserInfo(ctx context.Context,req*corpus.ListUserInfoReq) (map[string]interface{},map[string]interface{}){

	conds := map[string]interface{}{
		"is_deleted" : false,
	}
	limit := map[string]interface{}{
		"limit" : req.Limit,
		"offset" : req.Offset,
	}
	if req.Phone != ""{
		conds["phone"] = req.Phone
	}
	if req.EMail != ""{
		conds["e_mail"] = req.EMail
	}
	if req.UserName != ""{
		conds["user_name"] = req.UserName
	}

	return conds,limit
}