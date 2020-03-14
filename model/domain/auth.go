package domain

const TableAuth = "corpus_access"

var EmptyAuth = &AuthInfo{}

type AuthInfo struct {
	Id int64
	AuthCode string
	AuthDescription string
	ServiceName string
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m AuthInfo) TableName() string {
	return TableAuth
}