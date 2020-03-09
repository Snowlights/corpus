package model

import (
	"context"
	"database/sql"
	"github.com/Snowlights/corpus/common"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB    *sql.DB
	// TODO: remove
)

type DBTx interface {
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func Prepare(ctx context.Context) {
	fun := "model.Prepare -->"

	switch common.CurrEnv {
	case common.EnvTypeLocal:
		db, err := sql.Open("mysql", "root:woaininana12.@tcp(127.0.0.1:3306)/corpus")
		if err != nil {
			log.Panicf( "%s %s connect db err:%v", ctx,fun, err)
			return
		}
		DB = db

	default:
		log.Panicf( " %s %s unknown env:%v", ctx, fun, common.CurrEnv)
		return
	}
	log.Println(ctx, "%s succeeded", fun)
	return
}