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

func AddAuth(ctx context.Context, req *corpus.AddAuthReq) *corpus.AddAuthRes {
	fun := "Controller.AddAuth -->"
	res := &corpus.AddAuthRes{}

	//todo check cookie
	pass := cache.CheckSuperAdmin(ctx, req.Cookie)
	if !pass {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "权限不足，非超级管理员不可添加",
		}
		return res
	}

	data, conds, audit := toAddAuth(ctx, req)
	dataList, err := daoimpl.AuthDao.GetAuth(ctx, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	if len(dataList) > 0 {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "已存在的权限码",
		}
		return res
	}

	lastInsertId, err := daoimpl.AuthDao.AddAuth(ctx, data)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	auditLastInsertId, err := addAudit(ctx, audit)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d", ctx, fun, auditLastInsertId)

	log.Printf("%v %v success ,lastInsertId %d", ctx, fun, lastInsertId)

	return res
}

func toAddAuth(ctx context.Context, req *corpus.AddAuthReq) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"auth_code":        req.AuthCode,
		"auth_description": req.AuthDescription,
		"service_name":     req.ServiceName,
		"created_at":       now,
		"created_by":       req.Cookie,
		"updated_at":       now,
		"updated_by":       req.Cookie,
		"is_deleted":       false,
	}
	conds := map[string]interface{}{
		"auth_code":  req.AuthCode,
		"is_deleted": false,
	}
	audit := map[string]interface{}{
		"table_name": domain.EmptyAuth.TableName(),
		"history":    fmt.Sprintf("%s add auth %v time %d ", req.Cookie, req.AuthCode, now),
		"activity":   fmt.Sprintf("%s add auth", req.Cookie),
		"content":    fmt.Sprintf("%v", data),
		"created_at": now,
		"created_by": req.Cookie,
		"is_deleted": false,
	}

	return data, conds, audit
}

func UpdateAuth(ctx context.Context, req *corpus.UpdateAuthReq) *corpus.UpdateAuthRes {
	fun := "Controller.UpdateAuth -->"
	res := &corpus.UpdateAuthRes{}
	pass := cache.CheckSuperAdmin(ctx, req.Cookie)
	if !pass {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "权限不足，非超级管理员不可修改",
		}
		return res
	}
	err := daoimpl.AuthTxDao.UpdateAuthTx(ctx, req)
	if err != nil {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}
	log.Printf("%v %v success ", ctx, fun)
	return res
}

func DelAuth(ctx context.Context, req *corpus.DelAuthReq) *corpus.DelAuthRes {
	fun := "Controller.DelAuth -->"
	res := &corpus.DelAuthRes{}
	pass := cache.CheckSuperAdmin(ctx, req.Cookie)
	if !pass {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "权限不足，非超级管理员不可删除",
		}
		return res
	}
	err := daoimpl.AuthTxDao.DelAuthTx(ctx, req)
	if err != nil {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}
	log.Printf("%v %v success ", ctx, fun)
	return res
}

func ListAuth(ctx context.Context, req *corpus.ListAuthReq) *corpus.ListAuthRes {
	fun := "Controller.ListAuth -->"
	res := &corpus.ListAuthRes{}
	limit, conds := toListAuth(ctx, req)
	dataList, err := daoimpl.AuthDao.ListAuth(ctx, limit, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	total, err := daoimpl.AuthDao.CountAuth(ctx, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}
	data := formCorpusAuthInfoList(dataList)
	res.Data = &corpus.ListAuthData{
		Items:  data,
		Total:  total,
		Offset: req.Offset + int64(len(dataList)),
		More:   (req.Offset + int64(len(dataList))) < total,
	}

	return res
}

func formCorpusAuthInfoList(input []*domain.AuthInfo) []*corpus.AuthInfo {
	var output []*corpus.AuthInfo

	for _, item := range input {
		output = append(output, &corpus.AuthInfo{
			Id:              item.Id,
			AuthName:        item.AuthCode,
			AuthDescription: item.AuthDescription,
			ServiceName:     item.ServiceName,
		})
	}
	return output
}

func toListAuth(ctx context.Context, req *corpus.ListAuthReq) (map[string]interface{}, map[string]interface{}) {
	limit := map[string]interface{}{
		"limit":  req.Limit,
		"offset": req.Offset,
	}
	conds := map[string]interface{}{
		"is_deleted": false,
	}
	return limit, conds
}

func AddUserAuth(ctx context.Context, req *corpus.AddUserAuthReq) *corpus.AddUserAuthRes {
	fun := "Controller.AddUserAuth -->"
	res := &corpus.AddUserAuthRes{}
	data, conds, audit := toAddUserAuth(ctx, req)

	pass := cache.CheckIsAdmin(req.Cookie) || cache.CheckSuperAdmin(ctx, req.Cookie)
	if !pass {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "权限不足，非管理员不可添加",
		}
		return res
	}

	dataList, err := daoimpl.UserAuthDao.GetUserAuth(ctx, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	if len(dataList) > 0 {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "已经存在的记录",
		}
		return res
	}

	lastInsertId, err := daoimpl.UserAuthDao.AddUserAuth(ctx, data)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}
	auditLastInsertId, err := addAudit(ctx, audit)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d", ctx, fun, auditLastInsertId)

	log.Printf("%v %v success ,lastInsertId %d", ctx, fun, lastInsertId)

	return res
}

func toAddUserAuth(ctx context.Context, req *corpus.AddUserAuthReq) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"user_id":    req.UserId,
		"auth_code":  req.AuthCode,
		"created_at": now,
		"created_by": req.Cookie,
	}
	conds := map[string]interface{}{
		"user_id":    req.UserId,
		"auth_code":  req.AuthCode,
		"is_deleted": false,
	}
	audit := map[string]interface{}{
		"table_name": domain.EmptyUserAuth.TableName(),
		"history":    fmt.Sprintf("%s add user %v auth %v time %d ", req.Cookie, req.UserId, req.AuthCode, now),
		"activity":   fmt.Sprintf("%s add user auth", req.Cookie),
		"content":    fmt.Sprintf("%v", data),
		"created_at": now,
		"created_by": req.Cookie,
		"is_deleted": false,
	}
	return data, conds, audit
}

func DelUserAuth(ctx context.Context, req *corpus.DelUserAuthReq) *corpus.DelUserAuthRes {
	fun := "Controller.DelUserAuth --> "
	res := &corpus.DelUserAuthRes{}

	pass := cache.CheckIsAdmin(req.Cookie) || cache.CheckSuperAdmin(ctx, req.Cookie)
	if !pass {
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: "权限不足，非管理员不可删除",
		}
		return res
	}

	data, conds, audit := toDelUserAuth(ctx, req)
	rowsAffected, err := daoimpl.UserAuthDao.DelUserAuth(ctx, data, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}
	auditLastInsertId, err := addAudit(ctx, audit)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d", ctx, fun, auditLastInsertId)

	log.Printf("%v %v success ,rowsAffected %d", ctx, fun, rowsAffected)
	return res
}

func toDelUserAuth(ctx context.Context, req *corpus.DelUserAuthReq) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_deleted": true,
	}
	conds := map[string]interface{}{
		"user_id":   req.UserId,
		"auth_code": req.AuthCode,
	}
	audit := map[string]interface{}{
		"table_name": domain.EmptyUserAuth.TableName(),
		"history":    fmt.Sprintf("%s del user %v auth %v time %d ", req.Cookie, req.UserId, req.AuthCode, now),
		"activity":   fmt.Sprintf("%s del user auth", req.Cookie),
		"content":    fmt.Sprintf("%v", data),
		"created_at": now,
		"created_by": req.Cookie,
		"is_deleted": false,
	}
	return data, conds, audit
}

func ListUserAuth(ctx context.Context, req *corpus.ListUserAuthReq) *corpus.ListUserAuthRes {
	fun := "Controller.ListUserAuth -->"
	res := &corpus.ListUserAuthRes{}

	limit, conds := toListUserAuth(ctx, req)
	dataList, err := daoimpl.UserAuthDao.ListUserAuth(ctx, limit, conds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	if len(dataList) == 0 {
		res.Data = &corpus.ListUserAuthData{}
		return res
	}

	var codeList []string

	for _, item := range dataList {
		codeList = append(codeList, item.AuthCode)
	}

	authCodeConds := map[string]interface{}{
		"auth_code": codeList,
	}

	authDataList, err := daoimpl.AuthDao.ListAuth(ctx, limit, authCodeConds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	total, err := daoimpl.AuthDao.CountAuth(ctx, authCodeConds)
	if err != nil {
		log.Fatalf("%v %v error %v", ctx, fun, err)
		res.Errinfo = &corpus.ErrorInfo{
			Ret: -1,
			Msg: err.Error(),
		}
		return res
	}

	data := formCorpusAuthInfoList(authDataList)
	res.Data = &corpus.ListUserAuthData{
		Items:  data,
		Total:  total,
		Offset: req.Offset + int64(len(dataList)),
		More:   (req.Offset + int64(len(dataList))) < total,
	}

	return res
}

func toListUserAuth(ctx context.Context, req *corpus.ListUserAuthReq) (map[string]interface{}, map[string]interface{}) {
	limit := map[string]interface{}{
		"limit":  req.Limit,
		"offset": req.Offset,
	}
	conds := map[string]interface{}{
		"is_deleted": false,
	}
	if req.UserId != 0 {
		conds["user_id"] = req.UserId
	}

	return limit, conds
}
