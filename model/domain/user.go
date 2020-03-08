package domain


const TableUser = "corpus_user"

var EmptyUser = &UserInfo{}

type UserInfo struct {
	Id int64
	UserName string
	UserDescription string
	E_mail string
	UserPassword string
	Phone string
	Token string
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m UserInfo) TableName() string {
	return TableUser
}

