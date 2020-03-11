package domain

const TableAdminUser = "corpus_admin"

var EmptyAdminUser = &AdminUserInfo{}

type AdminUserInfo struct {
	Id int64
	UserId int64
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m AdminUserInfo) TableName() string {
	return TableAdminUser
}