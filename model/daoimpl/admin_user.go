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

func NewAdminUserDao() dao.AdminUserDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &AdminUserLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}
type AdminUserLocalImpl struct {
	DB *sql.DB
}

func  (m *AdminUserLocalImpl) GetAdminUser(ctx context.Context,conds map[string]interface{}) ([]*domain.AdminUserInfo,error){
	return getAdminUser(ctx,m.DB,conds)
}

func (m *AdminUserLocalImpl) AddAdminUser(ctx context.Context,data map[string]interface{}) (int64,error) {
	return addAdminUser(ctx,m.DB,data)
}

func (m *AdminUserLocalImpl) DelAdminUser(ctx context.Context,data,conds map[string]interface{}) (int64,error){
	return delAdminUser(ctx,m.DB,data,conds)
}

func (m *AdminUserLocalImpl) ListAdminUser(ctx context.Context,limit ,conds map[string]interface{}) ([]*domain.AdminUserInfo,error){
	return listAdminUser(ctx, m.DB,limit,conds)
}

func (m *AdminUserLocalImpl) CountAdminUser(ctx context.Context,conds map[string]interface{}) (int64,error){
	return countAdminUser(ctx,m.DB,conds)
}

func getAdminUser(ctx context.Context,db model.DBTx,conds map[string]interface{}) ([]*domain.AdminUserInfo,error){
	fun := "getAdminUser -->"

	cond := buildGet(conds,domain.EmptyAdminUser.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.AdminUserInfo

	for rows.Next(){
		var adminUserInfo domain.AdminUserInfo
		err = rows.Scan(&adminUserInfo.Id,&adminUserInfo.UserId,&adminUserInfo.CreatedAt,&adminUserInfo.CreatedBy,
			&adminUserInfo.UpdatedAt,&adminUserInfo.UpdatedBy,&adminUserInfo.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&adminUserInfo)
	}

	return res,nil
}

func addAdminUser(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addAdminUser -->"
	cond := buildInsert(data,domain.EmptyAdminUser.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func delAdminUser(ctx context.Context,db model.DBTx, data,conds map[string]interface{}) (int64,error){
	fun := "delAdminUser -->"

	cond := buildDelete(data,conds,domain.EmptyAdminUser.TableName())
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.RowsAffected()

}

func listAdminUser(ctx context.Context,db model.DBTx,limit, conds map[string]interface{}) ([]*domain.AdminUserInfo,error){
	fun:= "listAdminUser -->"

	cond := buildList(limit,conds,domain.EmptyAdminUser.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.AdminUserInfo

	for rows.Next(){
		var adminUserInfo domain.AdminUserInfo
		err = rows.Scan(&adminUserInfo.Id,&adminUserInfo.UserId,&adminUserInfo.CreatedAt,&adminUserInfo.CreatedBy,
			&adminUserInfo.UpdatedAt,&adminUserInfo.UpdatedBy,&adminUserInfo.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&adminUserInfo)
	}

	return res,nil
}

func countAdminUser(ctx context.Context,db model.DBTx,conds map[string]interface{}) (total int64,err error){
	fun:= "countAdminUser -->"

	cond := buildCount(conds,domain.EmptyAdminUser.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)

	return
}