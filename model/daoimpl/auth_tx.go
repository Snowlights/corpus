package daoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Snowlights/corpus/common"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/dao"
	"github.com/Snowlights/corpus/model/domain"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"time"
)

func NewAuthTxDao() dao.AuthTxDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &AuthTxLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type AuthTxLocalImpl struct {
	DB *sql.DB
}

func (m *AuthTxLocalImpl) UpdateAuthTx(ctx context.Context,req *corpus.UpdateAuthReq) error{
	return updateAuthTx(ctx,m.DB,req)
}

func (m *AuthTxLocalImpl) DelAuthTx(ctx context.Context,req *corpus.DelAuthReq) error{
	return delAuthTx(ctx,m.DB,req)
}

func updateAuthTx(ctx context.Context,db model.DBTx,req *corpus.UpdateAuthReq) error{
	fun := "updateAuthTx -->"
	sdb, ok := db.(*sql.DB)
	if !ok {
		return fmt.Errorf("error sql tx")
	}
	tx , err := sdb.Begin()
	if err != nil{
		return err
	}

	limit := map[string]interface{}{
		"offset" : 0,
		"limit" : 1,
	}
	conds := map[string]interface{}{
		"id" : req.Id,
		"is_deleted" : false,
	}
	authInfo,err := listAuth(ctx,tx,limit,conds)
	if err != nil{
		tx.Rollback()
		return err
	}
	if len(authInfo) == 0 {
		tx.Rollback()
		return fmt.Errorf("未找到该权限信息")
	}

	updateConds := map[string]interface{}{
		"auth_code" : authInfo[0].AuthCode,
	}
	updateData := map[string]interface{}{
		"auth_code" : req.AuthCode,
	}
	data,cond,audit := toUpdateAuth(ctx,req)
	authRowsAffected,err := updateAuth(ctx,tx,data,cond)
	if err != nil{
		tx.Rollback()
		return fmt.Errorf("修改权限码失败")
	}
	userAuthRowsAffected,err := updateUserAuth(ctx,tx,updateData,updateConds)
	if err != nil{
		tx.Rollback()
		return fmt.Errorf("修改映射表失败")
	}
	log.Printf("%v %v authRowsAffected %d userAuthRowsAffected %d",ctx,fun,authRowsAffected,userAuthRowsAffected)
	auditLastInsertId, err := addAudit(ctx,tx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)

	return tx.Commit()
}

func toUpdateAuth(ctx context.Context,req *corpus.UpdateAuthReq) (map[string]interface{},map[string]interface{},map[string]interface{}){
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
	audit := map[string]interface{}{
		"table_name" : domain.EmptyAuth.TableName(),
		"history" : fmt.Sprintf("%s update auth %v time %d ",req.Cookie,req.AuthCode,now),
		"activity" : fmt.Sprintf("%s update auth",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,conds,audit
}


func delAuthTx(ctx context.Context,db model.DBTx,req *corpus.DelAuthReq) error{
	fun := "updateAuthTx -->"
	sdb, ok := db.(*sql.DB)
	if !ok {
		return fmt.Errorf("error sql tx")
	}
	tx , err := sdb.Begin()
	if err != nil{
		return err
	}

	limit := map[string]interface{}{
		"offset" : 0,
		"limit" : 1,
	}
	conds := map[string]interface{}{
		"is_deleted" : false,
	}
	updateData := map[string]interface{}{
		"is_deleted" : true,
	}
	updateConds := make(map[string]interface{})
	if req.Id != 0{
		conds["id"] = req.Id
		authInfo,err := listAuth(ctx,tx,limit,conds)
		if err != nil{
			tx.Rollback()
			return fmt.Errorf("未找到该权限信息")
		}
		updateConds["auth_code"] = authInfo[0].AuthCode
	} else {
		updateConds["auth_code"] = req.AuthCode
	}
	data,cond,audit := toDelAuth(ctx,req)
	authRowsAffected,err := delAuth(ctx,tx,data,cond)
	if err != nil{
		tx.Rollback()
		return fmt.Errorf("删除权限失败")
	}
	userAuthRowsAffected,err := delUserAuth(ctx,tx,updateData,updateConds)
	if err != nil{
		tx.Rollback()
		return fmt.Errorf("修改映射表失败")
	}
	log.Printf("%v %v authRowsAffected %d userAuthRowsAffected %d",ctx,fun,authRowsAffected,userAuthRowsAffected)
	auditLastInsertId, err := addAudit(ctx,tx,audit)
	if err!= nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}
	log.Printf("%v %v success ,auditLastInsertId %d",ctx,fun,auditLastInsertId)
	return tx.Commit()
}

func toDelAuth(ctx context.Context,req* corpus.DelAuthReq) (map[string]interface{},map[string]interface{},map[string]interface{}){
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
	audit := map[string]interface{}{
		"table_name" : domain.EmptyAuth.TableName(),
		"history" : fmt.Sprintf("%s del auth %v time %d ",req.Cookie,req.AuthCode,now),
		"activity" : fmt.Sprintf("%s del auth",req.Cookie),
		"content" : fmt.Sprintf("%v",data),
		"created_at" : now,
		"created_by" : req.Cookie,
		"is_deleted" : false,
	}
	return data,conds,audit
}