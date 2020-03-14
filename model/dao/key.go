package dao

import "context"

type KeyDao interface {
	AddKey(ctx context.Context,data []map[string]interface{}) (int64,error)
}