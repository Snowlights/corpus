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

func NewEvaluationDao() dao.EvaluationDao {
	fun := "NewUserDao -->"
	switch common.CurrEnv {
	case common.EnvTypeLocal:
		return &EvaluationLocalImpl{model.DB}
	default:
		log.Panic(context.Background(), "unknown env:%v", fun, common.CurrEnv)
	}
	return nil
}

type EvaluationLocalImpl struct {
	DB *sql.DB
}

func (m *EvaluationLocalImpl) AddEvaluation(ctx context.Context, data map[string]interface{}) (int64, error) {
	return addEvaluation(ctx, m.DB, data)
}
func (m *EvaluationLocalImpl) ListEvaluation(ctx context.Context, limit, conds map[string]interface{}) ([]*domain.EvaluationInfo, error) {
	return listEvaluation(ctx,m.DB,limit,conds)
}
func (m *EvaluationLocalImpl) CountEvaluation(ctx context.Context, conds map[string]interface{}) (int64, error) {
	return countEvaluation(ctx,m.DB,conds)
}

func addEvaluation(ctx context.Context,db model.DBTx,data map[string]interface{})(int64,error){
	fun := "addEvaluation -->"

	cond := buildInsert(data,domain.EmptyEvaluation.TableName())
	log.Printf("%v",cond)
	sqlResult, err := db.Exec(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func listEvaluation(ctx context.Context,db model.DBTx,limit,conds map[string]interface{}) ([]*domain.EvaluationInfo, error){
	fun := "listEvaluation -->"

	cond := buildList(limit,conds,domain.EmptyEvaluation.TableName())

	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}

	var res []*domain.EvaluationInfo

	for rows.Next(){
		var evaluation domain.EvaluationInfo
		err = rows.Scan(&evaluation.Id,&evaluation.AudioSrc,&evaluation.AudioText,&evaluation.TotalScore,&evaluation.OriginalData,
			&evaluation.CreatedAt,&evaluation.CreatedBy,&evaluation.UpdatedAt,&evaluation.UpdatedBy,&evaluation.IsDeleted)
		res = append(res,&evaluation)
	}

	return res,nil
}

func countEvaluation(ctx context.Context,db model.DBTx,conds map[string]interface{}) (total int64,err error){

	fun:= "countEvaluation -->"

	cond := buildCount(conds,domain.EmptyEvaluation.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)

	return
}