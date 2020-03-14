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

func NewAuthrDao() dao.AuthDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &AuthLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}
type AuthLocalImpl struct {
	DB *sql.DB
}

func (m *AuthLocalImpl)GetAuth(ctx context.Context,conds map[string]interface{}) ([]*domain.AuthInfo,error){
	return getAuth(ctx,m.DB,conds)
}
func (m *AuthLocalImpl) AddAuth(ctx context.Context,data map[string]interface{}) (int64,error){
	return addAuth(ctx,m.DB,data)
}
func (m *AuthLocalImpl) UpdateAuth(ctx context.Context,data,conds map[string]interface{}) (int64,error){
	return updateAuth(ctx,m.DB,data,conds)
}
func (m *AuthLocalImpl) DelAuth(ctx context.Context,data, conds map[string]interface{}) (int64,error){
	return delAuth(ctx,m.DB,data,conds)
}
func (m *AuthLocalImpl) ListAuth(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.AuthInfo,error){
	return listAuth(ctx,m.DB,limit,conds)
}
func (m *AuthLocalImpl) CountAuth(ctx context.Context,conds map[string]interface{}) (int64,error){
	return countAuth(ctx,m.DB,conds)
}

func getAuth(ctx context.Context,db model.DBTx,conds map[string]interface{}) ([]*domain.AuthInfo,error){
	fun := "getAuth -->"

	cond := buildGet(conds,domain.EmptyAuth.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}
	var res []*domain.AuthInfo
	for rows.Next(){
		var auth domain.AuthInfo
		err = rows.Scan(&auth.Id,&auth.AuthCode,&auth.AuthDescription,&auth.ServiceName,&auth.CreatedAt,
			&auth.CreatedBy,&auth.UpdatedAt,&auth.UpdatedBy,&auth.IsDeleted)
		res = append(res,&auth)
	}

	return res,nil
}

func addAuth(ctx context.Context,db model.DBTx, data map[string]interface{}) (int64,error){
	fun := "addAuth -->"
	cond := buildInsert(data,domain.EmptyAuth.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func updateAuth(ctx context.Context,db model.DBTx,data,conds map[string]interface{}) (int64,error){
	fun := "updateAuth -->"
	cond := buildUpdate(data,conds,domain.EmptyAuth.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}

	return sqlResult.RowsAffected()
}
func delAuth(ctx context.Context,db model.DBTx,data, conds map[string]interface{}) (int64,error){
	fun := "delAuth -->"

	cond := buildDelete(data,conds,domain.EmptyAuth.TableName())
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.RowsAffected()
}
func listAuth(ctx context.Context,db model.DBTx, limit,conds map[string]interface{}) ([]*domain.AuthInfo,error) {
	fun := "listAuth -->"
	cond := buildList(limit,conds,domain.EmptyAuth.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}
	var res []*domain.AuthInfo
	for rows.Next(){
		var auth domain.AuthInfo
		err = rows.Scan(&auth.Id,&auth.AuthCode,&auth.AuthDescription,&auth.ServiceName,&auth.CreatedAt,
			&auth.CreatedBy,&auth.UpdatedAt,&auth.UpdatedBy,&auth.IsDeleted)
		res = append(res,&auth)
	}

	return res,nil
}
func countAuth(ctx context.Context,db model.DBTx,conds map[string]interface{}) (total int64,err error){
	fun:= "countAuth -->"

	cond := buildCount(conds,domain.EmptyAuth.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)

	return
}