package daoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Snowlights/corpus/common"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/dao"
	"log"
)

func NewKeyWordTxDao() dao.KeyWordTxDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &KeyWordTxLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type KeyWordTxLocalImpl struct {
	DB *sql.DB
}

func (m* KeyWordTxLocalImpl) AddKeyWordTx(ctx context.Context,keyWordData map[string]interface{},keyData []map[string]interface{}) (error){

	return addKeyWordTx(ctx,m.DB,keyWordData,keyData)
}

func addKeyWordTx (ctx context.Context,db model.DBTx,keyWordData map[string]interface{},keyData []map[string]interface{}) error{
	fun := "addKeyWordTx -->"
	sdb, ok := db.(*sql.DB)
	if !ok {
		return fmt.Errorf("error sql tx")
	}
	tx , err := sdb.Begin()
	if err != nil{
		return err
	}

	keyWordLastInsertId,err := addKeyWord(ctx,tx,keyWordData)
	if err != nil{
		tx.Rollback()
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return err
	}
	log.Printf("%v %v keyWordLastInsertId %d ",ctx,fun,keyWordLastInsertId)

	for _, item := range keyData{
		item["keyword_id"] = keyWordLastInsertId
	}

	keyLastInsertId ,err := addKey(ctx,tx,keyData)
	if err != nil{
		tx.Rollback()
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return err
	}
	log.Printf("%v %v keyLastInsertId %d ",ctx,fun,keyLastInsertId)

	return tx.Commit()
}