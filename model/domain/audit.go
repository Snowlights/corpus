package domain

const TableAudit = "corpus_audit"

var EmptyAudit = &AuditInfo{}

type AuditInfo struct {
	Id int64
	UserId int64
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m AuditInfo) TableName() string {
	return TableAudit
}