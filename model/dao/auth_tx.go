package dao

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
)

type AuthTxDao interface {
	UpdateAuthTx(ctx context.Context,req *corpus.UpdateAuthReq) error
	DelAuthTx(ctx context.Context,req *corpus.DelAuthReq) error
}
