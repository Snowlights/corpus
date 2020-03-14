package dao

import "context"

type KeyWordTxDao interface {
	AddKeyWordTx(ctx context.Context,keyWordData map[string]interface{},keyData []map[string]interface{}) (error)
}