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

func NewAuditDao() dao.AuditDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &AuditLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}
type AuditLocalImpl struct {
	DB *sql.DB
}

func (m *AuditLocalImpl) AddAudit(ctx context.Context,data map[string]interface{}) (int64,error) {
	return addAudit(ctx,m.DB,data)
}

func addAudit(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addAudit -->"

	cond := buildInsert(data,domain.EmptyAudit.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}