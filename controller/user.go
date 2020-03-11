package controller

import (
	"context"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"time"
)

func LoginUser(ctx context.Context,req *corpus.LoginUserReq) *corpus.LoginUserRes{
	fun := "Controller.LoginUser -- >"
	res := &corpus.LoginUserRes{}

	data,conds := toLoginUser(ctx,req)

	dataList ,err := daoimpl.UserDao.GetUserInfo(ctx,conds)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  err.Error(),
		}
		return res
	}
	if len(dataList) > 0 {
		if dataList[0].UserPassword != req.UserPassword{
			res.Errinfo = &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  "密码不正确",
			}
			return res
		}
		AddCookieToList(dataList[0].Token)
		if req.Phone != ""{
			DelPhoneCode(req.Phone)
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
		SendEmail([]string{req.EMail})
	}

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return res
}

func LoginOutUserInfo(ctx context.Context, req *corpus.LogoutUserInfoReq) *corpus.LogoutUserInfoRes{
	fun := "Controller.LoginOutUserInfo -- >"
	res := &corpus.LogoutUserInfoRes{}

	pass := CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}
	pass = DelCookieFromList(req.Cookie)
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

	pass := CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}

	conds, data := toUpdateUserInfo(ctx,req)

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

	pass := CheckOnLine(req.Cookie)
	if !pass{
		log.Fatalf("%v %s err cookie not in list %v",ctx,fun,req.Cookie)
		res.Errinfo = &corpus.ErrorInfo{
			Ret:                  -1,
			Msg:                  "未检测到登陆信息",
		}
		return res
	}
	conds, data := toDelUserInfo(ctx,req)
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

func toLoginUser(ctx context.Context, req *corpus.LoginUserReq) (map[string]interface{}, map[string]interface{}){
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
	if req.Phone != ""{
		data["user_name"] = req.Phone
		data["phone"] = req.Phone
		conds["phone"] = req.Phone
	}

	if req.EMail != ""{
		data["user_name"] = req.EMail
		data["e_mail"] = req.EMail
		data["user_password"] = req.UserPassword
		conds["e_mail"] = req.EMail
	}
	return data,conds
}

func toUpdateUserInfo(ctx context.Context,req *corpus.UpdateUserInfoReq) (map[string]interface{},map[string]interface{}){
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

	return conds,data
}

func toDelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq) (map[string]interface{},map[string]interface{}){
	now := time.Now().Unix()
	conds := map[string]interface{}{
		"id" : req.UserId,
	}
	data := map[string]interface{}{
		"is_deleted" : true,
		"updated_at" : now,
		"updated_by" : "admin",
	}
	return conds,data
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