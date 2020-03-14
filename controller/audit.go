package controller

import (
	"context"
	"github.com/Snowlights/corpus/model/daoimpl"
	"log"
)

func addAudit(ctx context.Context,data map[string]interface{}) (int64,error){
	fun := "Controller.addAudit -->"

	lastInsertId, err := daoimpl.AuditDao.AddAudit(ctx,data)
	if err != nil{
		log.Fatalf("%v %v error %v",ctx,fun,err)
	}

	log.Printf("%v %v success ,lastInsertId %d",ctx,fun,lastInsertId)
	return lastInsertId,err
}