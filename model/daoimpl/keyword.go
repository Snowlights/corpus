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

func NewKeyWordDao() dao.KeyWordDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &KeyWordLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type KeyWordLocalImpl struct {
	DB *sql.DB
}

func (m* KeyWordLocalImpl) AddKeyWord(ctx context.Context,data map[string]interface{}) (int64,error){
	return addKeyWord(ctx,m.DB,data)
}


func addKeyWord(ctx context.Context,db model.DBTx, data map[string]interface{}) (int64,error){
	fun := "addKeyWord -->"

	cond := buildInsert(data,domain.EmptyKeyWord.TableName())
	log.Printf("%v",cond)
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}