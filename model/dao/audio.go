package dao

import (
	"context"
	"github.com/Snowlights/corpus/model/domain"
)

type AudioDao interface {
	AddAudio(ctx context.Context,data map[string]interface{}) (int64,error)
	DelAudio(ctx context.Context,data,conds map[string]interface{}) (int64,error)
	ListAudio(ctx context.Context,limit,conds map[string]interface{}) ([]*domain.AudioInfo,error)
	CountAudio(ctx context.Context,conds map[string]interface{}) (int64,error)
	AddAudioToText(ctx context.Context,data map[string]interface{}) (int64,error)
}
