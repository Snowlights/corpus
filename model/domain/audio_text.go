package domain
const TableAudioText = "corpus_audio_text"

var EmptyAudioText = &AudioTextInfo{}

type AudioTextInfo struct {
	Id int64
	AudioSrc string
	AudioTransFrom string
	AudioText string
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m AudioTextInfo) TableName() string {
	return TableAudioText
}
