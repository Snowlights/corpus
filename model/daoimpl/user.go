package daoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Snowlights/corpus/common"
	"github.com/Snowlights/corpus/model/dao"
	"log"
)

func NewCommentDao() dao.UserDao {
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

func (m UserLocalImpl) AddUserInfo(ctx context.Context,dataList map[string]interface{}) (int64, error){
	return addUserInfo(ctx,m.DB,dataList)
}

func addUserInfo(ctx context.Context,db model.DBTx,dataList map[string]interface{}) (int64,error){
	fun := "addUserInfo -->"

	cond := fmt.Sprintf("insert into %s values (?,?,?,?,?,?,?,?,?,?,?)",)
	

}
