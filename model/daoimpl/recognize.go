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
func (m* RecognizeLocalImpl) ListImageByCookie(ctx context.Context,limit,conds map[string]interface{}) ([]string,error){
	return listImageByCookie(ctx,m.DB,limit,conds)
}
func (m *RecognizeLocalImpl) CountImageByCookie(ctx context.Context,conds map[string]interface{})(int64,error){
	return countImageByCookie(ctx,m.DB,conds)
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

func listImageByCookie(ctx context.Context,db model.DBTx,limit,conds map[string]interface{}) ([]string,error){
	fun := "listImageBuCookie -->"
	cond := buildList(limit,conds,domain.EmptyPicture.TableName())

	rows,err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return nil, err
	}
	var res []string

	for rows.Next(){
		var picture domain.PictureInfo
		err = rows.Scan(&picture.Id,&picture.PictureSrc,&picture.PictureDes,&picture.Md5,&picture.PictureText,
			&picture.CreatedAt,&picture.CreatedBy,&picture.IsDeleted)
		res = append(res,picture.Md5)
	}
	err = rows.Err()
	return res,err
}

func countImageByCookie(ctx context.Context,db model.DBTx,conds map[string]interface{}) (total int64,err error){
	fun := "countUserInfo -->"
	cond := buildCount(conds,domain.EmptyPicture.TableName())
	rows, err := db.Query(cond)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&total)
	if err != nil {
		return
	}
	return
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