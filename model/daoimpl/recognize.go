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

func NewRecognizeDao() dao.RecognizeDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &RecognizeLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type RecognizeLocalImpl struct {
	DB *sql.DB
}
func (m* RecognizeLocalImpl) AddRecognizeImage(ctx context.Context,data map[string]interface{}) (int64,error){
	return addRecognizeImage(ctx,m.DB,data)
}
func (m* RecognizeLocalImpl) AddRecognizeAge(ctx context.Context,data map[string]interface{}) (int64,error){
	return addRecognizeAge(ctx,m.DB,data)
}
func addRecognizeImage(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addRecognizeImage -->"
	cond := buildInsert(data,domain.EmptyPicture.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func addRecognizeAge(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addRecognizeAge -->"
	cond := buildInsert(data,domain.EmptyRecognize.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}