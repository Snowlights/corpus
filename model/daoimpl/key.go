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

func NewKeyDao() dao.KeyDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &KeyLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type KeyLocalImpl struct {
	DB *sql.DB
}

func (m *KeyLocalImpl) AddKey(ctx context.Context,data []map[string]interface{}) (int64,error){
	return addKey(ctx,m.DB,data)
}

func addKey(ctx context.Context,db model.DBTx,data []map[string]interface{}) (int64,error){
	fun := "addKey -->"
	i := 0
	log.Printf("%d",len(data))
	for i < len(data){
		cond := buildInsert(data[i],domain.EmptyKey.TableName())
		_, err := db.Exec(cond)
		if err != nil{
			log.Fatalf("%v %v error %v",ctx,fun,err)
			return 0, err
		}
		i++
	}

	return 0,nil
}

