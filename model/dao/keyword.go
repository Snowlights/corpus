package dao

import "context"

type KeyWordDao interface {
	AddKeyWord(ctx context.Context,data map[string]interface{}) (int64,error)
}