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

func (m UserLocalImpl) ListUserInfo(ctx context.Context,conds map[string]interface{}) ([]*domain.UserInfo,error){
	return listUserInfo(ctx,m.DB,conds)
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
		err = rows.Scan(&userInfo.Id,&userInfo.UserName,&userInfo.UserDescription,&userInfo.E_mail,&userInfo.UserPassword,
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

	cond := fmt.Sprintf("insert into %s values (%s,%s,%s,%s,%s,%s,%d,%s,%d,%s,%d)",domain.EmptyUser.TableName(),
		dataList["user_name"],dataList["user_description"],dataList["e_mail"],dataList["user_password"],dataList["phone"],
		dataList["token"],dataList["created_at"],dataList["created_by"],dataList["updated_at"],dataList["updated_by"],dataList["is_deleted"])

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func updateUserInfo(ctx context.Context, db model.DBTx,data,conds map[string]interface{}) (int64,error){

}

func delUserInfo(ctx context.Context,db model.DBTx,data,conds map[string]interface{}) (int64,error){

}

func listUserInfo(ctx context.Context,db model.DBTx,conds map[string]interface{}) ([]*domain.UserInfo,error){

}