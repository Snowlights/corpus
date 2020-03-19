package daoimpl

import (
	"context"
	"github.com/Snowlights/corpus/model/dao"
	"log"
)

var (
	UserDao dao.UserDao
	AdminUserDao dao.AdminUserDao
	AuthDao dao.AuthDao
	UserAuthDao dao.UserAuthDao
	AuditDao dao.AuditDao
	EvaluationDao dao.EvaluationDao
	KeyDao dao.KeyDao
	KeyWordDao dao.KeyWordDao
	KeyWordTxDao dao.KeyWordTxDao
	AudioDao dao.AudioDao
	RecognizeDao dao.RecognizeDao
)

func PrePare(ctx context.Context){
	fun := "GrpcController.Prepare -->"

	UserDao = NewUserDao()
	AdminUserDao = NewAdminUserDao()
	AuthDao = NewAuthrDao()
	UserAuthDao = NewUserAuthDao()
	AuditDao = NewAuditDao()
	EvaluationDao = NewEvaluationDao()
	KeyDao = NewKeyDao()
	KeyWordDao = NewKeyWordDao()
	KeyWordTxDao = NewKeyWordTxDao()
	AudioDao = NewAudioDao()
	RecognizeDao = NewRecognizeDao()
	log.Printf("%v %s success ",ctx,fun)
	return
}