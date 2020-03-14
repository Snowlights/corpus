package domain

const TableKey = "corpus_key"

var EmptyKey = &KeyInfo{}

type KeyInfo struct {
	Id int64
	KeyWordId int64
	Word string
	Score string
	CreatedAt int64
	CreatedBy string
	IsDeleted bool
}
func (m KeyInfo) TableName() string {
	return TableKey
}
