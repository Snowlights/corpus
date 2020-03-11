package daoimpl

import (
	"context"
	"database/sql"
	"github.com/Snowlights/corpus/common"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/dao"
	"github.com/Snowlights/corpus/model/domain"
	"log"
)

func NewUserAuthDao() dao.UserAuthDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &UserAuthLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}
type UserAuthLocalImpl struct {
	DB *sql.DB
}
func (m *UserAuthLocalImpl) GetUserAuth(ctx context.Context,conds map[string]interface{})([]*domain.UserAuthInfo,error){
	return getUserAuth(ctx,m.DB,conds)
}
func (m *UserAuthLocalImpl) AddUserAuth(ctx context.Context,data map[string]interface{}) (int64,error){
	return addUserAuth(ctx,m.DB,data)
}
func (m *UserAuthLocalImpl) DelUserAuth(ctx context.Context,data,conds map[string]interface{}) (int64,error){
	return delUserAuth(ctx,m.DB,data,conds)
}
func (m *UserAuthLocalImpl) ListUserAuth(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.UserAuthInfo,error){
	return listUserAuth(ctx,m.DB,limit,conds)
}
func (m *UserAuthLocalImpl) CountUserAuth(ctx context.Context,conds map[string]interface{}) (int64,error){
	return countUserAuth(ctx,m.DB,conds)
}

func getUserAuth(ctx context.Context,db model.DBTx,conds map[string]interface{})([]*domain.UserAuthInfo,error){
	fun := "getUserAuth -->"
	cond := buildGet(conds,domain.EmptyUser.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}
	var res []*domain.UserAuthInfo
	for rows.Next(){
		var userAuth domain.UserAuthInfo
		err = rows.Scan(&userAuth.Id,&userAuth.UserId,&userAuth.AuthCode,&userAuth.CreatedAt,&userAuth.CreatedBy,&userAuth.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&userAuth)
	}

	return res,nil
}


func addUserAuth(ctx context.Context,db model.DBTx, data map[string]interface{}) (int64,error){
	fun := "addUserAuth --> "

	cond := buildInsert(data,domain.EmptyUserAuth.TableName())
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}
func delUserAuth(ctx context.Context,db model.DBTx,data,conds map[string]interface{}) (int64,error){
	fun := "delUserAuth -->"
	cond := buildDelete(data,conds,domain.EmptyUser.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}

	return sqlResult.RowsAffected()
	return 0,nil
}
func listUserAuth(ctx context.Context,db model.DBTx,limit,conds map[string]interface{}) ([]*domain.UserAuthInfo,error){
	fun := "listUserAuth -->"
	cond := buildList(limit,conds,domain.EmptyUser.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.UserAuthInfo

	for rows.Next(){
		var userAuth domain.UserAuthInfo
		err = rows.Scan(&userAuth.Id,&userAuth.UserId,&userAuth.AuthCode,&userAuth.CreatedAt,&userAuth.CreatedBy,&userAuth.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&userAuth)
	}

	return res,nil
}

func countUserAuth(ctx context.Context,db model.DBTx,conds map[string]interface{}) (total int64,err error){
	fun := "countUserAuth -->"
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