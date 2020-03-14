package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)

type EvaluationDao interface {
	AddEvaluation(ctx context.Context,data map[string]interface{}) (int64,error)
	ListEvaluation(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.EvaluationInfo,error)
	CountEvaluation(ctx context.Context,conds map[string]interface{}) (int64,error)
}