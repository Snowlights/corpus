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

func NewAudioDao() dao.AudioDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &AudioLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}
type AudioLocalImpl struct {
	DB *sql.DB
}

func (m *AudioLocalImpl ) AddAudio(ctx context.Context,data map[string]interface{}) (int64,error){
	return addAudio(ctx,m.DB,data)
}
func (m *AudioLocalImpl ) DelAudio(ctx context.Context,data,conds map[string]interface{}) (int64,error){
	return delAudio(ctx,m.DB,data,conds)
}
func (m *AudioLocalImpl ) ListAudio(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.AudioInfo,error){
	return listAudio(ctx,m.DB,limit,conds)
}
func (m *AudioLocalImpl ) CountAudio(ctx context.Context,conds map[string]interface{}) (int64,error){
	return countAudio(ctx,m.DB,conds)
}
func (m *AudioLocalImpl ) AddAudioToText(ctx context.Context,data map[string]interface{}) (int64,error){
	return addAudioToText(ctx,m.DB,data)
}

func addAudio(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addAudio -->"
	cond := buildInsert(data,domain.EmptyAudio.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func delAudio(ctx context.Context,db model.DBTx,data,conds map[string]interface{}) (int64,error){
	fun := "DelAudio -->"

	cond := buildDelete(data,conds,domain.EmptyAudio.TableName())
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.RowsAffected()
}

func listAudio(ctx context.Context,db model.DBTx,limit,conds map[string]interface{}) ([]*domain.AudioInfo,error){
	fun:= "ListAudio -->"

	cond := buildList(limit,conds,domain.EmptyAudio.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.AudioInfo

	for rows.Next(){
		var audio domain.AudioInfo
		err = rows.Scan(&audio.Id,&audio.AudioSrc,&audio.AudioDes,&audio.AudioType,&audio.CreatedAt,&audio.CreatedBy,
			&audio.UpdatedAt,&audio.UpdatedBy,&audio.IsDeleted)
		if err != nil{
			return nil,err
		}
		res = append(res,&audio)
	}

	return res,nil
}

func countAudio(ctx context.Context,db model.DBTx, conds map[string]interface{}) (total int64,err error){
	fun:= "CountAudio -->"

	cond := buildCount(conds,domain.EmptyAudio.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)

	return
}

func addAudioToText(ctx context.Context,db model.DBTx,data map[string]interface{}) (int64,error){
	fun := "addAudioToText -->"
	cond := buildInsert(data,domain.EmptyAudioText.TableName())

	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}