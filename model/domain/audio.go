package domain

import corpus "github.com/Snowlights/pub/grpc"

const TableAudio = "corpus_audio"

var EmptyAudio = &AudioInfo{}

type AudioInfo struct {
	Id int64
	AudioSrc string
	AudioDes string
	AudioType corpus.AudioType
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m AudioInfo) TableName() string {
	return TableAudio
}
