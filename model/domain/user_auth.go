package domain

const TableUserAuth = "corpus_user_access"

var EmptyUserAuth = &UserAuthInfo{}

type UserAuthInfo struct {
	Id int64
	UserId int64
	AuthCode string
	CreatedAt int64
	CreatedBy string
	IsDeleted bool
}
func (m UserAuthInfo) TableName() string {
	return TableUserAuth
}
