package domain
const TableKeyWord = "corpus_keyword"

var EmptyKeyWord= &KeyWordInfo{}

type KeyWordInfo struct {
	Id int64
	OriginTest string
	CreatedAt int64
	CreatedBy string
	IsDeleted bool
}
func (m KeyWordInfo) TableName() string {
	return TableKeyWord
}
