package dao

import "context"

type AuditDao interface {
	AddAudit(ctx context.Context,data map[string]interface{}) (int64,error)
}
