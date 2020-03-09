package daoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Snowlights/corpus/common"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/dao"
	"github.com/Snowlights/corpus/model/domain"
	"log"
)

func NewUserDao() dao.UserDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &UserLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type UserLocalImpl struct {
	DB *sql.DB
}

func (m UserLocalImpl) GetUserInfo(ctx context.Context,conds map[string]interface{}) ([]*domain.UserInfo,error){
	return getUserInfo(ctx,m.DB,conds)
}

func (m UserLocalImpl) AddUserInfo(ctx context.Context,dataList map[string]interface{}) (int64, error){
	return addUserInfo(ctx,m.DB,dataList)
}

func (m UserLocalImpl) UpdateUserInfo(ctx context.Context,data,conds map[string]interface{}) (int64,error){
	return updateUserInfo(ctx,m.DB,data,conds)
}

func (m UserLocalImpl) DelUserInfo(ctx context.Context, data,conds map[string]interface{})(int64,error){
	return delUserInfo(ctx,m.DB,data,conds)
}

func (m UserLocalImpl) ListUserInfo(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.UserInfo,error){
	return listUserInfo(ctx,m.DB,limit,conds)
}

func (m UserLocalImpl) CountUserInfo(ctx context.Context,conds map[string]interface{}) (int64,error){
	return countUserInfo(ctx,m.DB,conds)
}

func getUserInfo(ctx context.Context,db model.DBTx,conds map[string]interface{}) ([]*domain.UserInfo,error){
	fun := "getUserInfo -->"

	cond := fmt.Sprintf("select * from %s where phone = %s or e_mail = %s limit 1",domain.EmptyUser.TableName(),conds["phone"],
		conds["e_mail"])

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.UserInfo

	for rows.Next(){
		var userInfo domain.UserInfo
		err = rows.Scan(&userInfo.Id,&userInfo.UserName,&userInfo.UserDescription,&userInfo.E_mail,&userInfo.Phone,&userInfo.UserPassword,
			&userInfo.Token,&userInfo.CreatedAt,&userInfo.CreatedBy,&userInfo.UpdatedAt,&userInfo.UpdatedBy,&userInfo.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&userInfo)
	}

	return res,nil
}

func addUserInfo(ctx context.Context,db model.DBTx,dataList map[string]interface{}) (int64,error){
	fun := "addUserInfo -->"

	cond := buildInsert(dataList,domain.EmptyUser.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func updateUserInfo(ctx context.Context, db model.DBTx,data,conds map[string]interface{}) (int64,error){

	fun := "updateUserInfo -->"
	cond := buildUpdate(data,conds,domain.EmptyUser.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}

	return sqlResult.RowsAffected()
}

func delUserInfo(ctx context.Context,db model.DBTx,data,conds map[string]interface{}) (int64,error){
	fun := "delUserInfo -->"
	cond := buildDelete(data,conds,domain.EmptyUser.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}

	return sqlResult.RowsAffected()
}

func listUserInfo(ctx context.Context,db model.DBTx,limit, conds map[string]interface{}) ([]*domain.UserInfo,error){

	fun := "delUserInfo -->"
	cond := buildList(limit,conds,domain.EmptyUser.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.UserInfo

	for rows.Next(){
		var userInfo domain.UserInfo
		err = rows.Scan(&userInfo.Id,&userInfo.UserName,&userInfo.UserDescription,&userInfo.E_mail,&userInfo.Phone,&userInfo.UserPassword,
			&userInfo.Token,&userInfo.CreatedAt,&userInfo.CreatedBy,&userInfo.UpdatedAt,&userInfo.UpdatedBy,&userInfo.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&userInfo)
	}

	return res,nil
}

func countUserInfo(ctx context.Context,db model.DBTx, conds map[string]interface{}) (total int64,err error) {
	fun := "countUserInfo -->"
	cond := buildCount(conds,domain.EmptyUser.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)
	if err != nil {
		return
	}

	return
}