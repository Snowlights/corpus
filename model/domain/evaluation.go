package domain
const TableEvaluation = "corpus_evaluation"

var EmptyEvaluation= &EvaluationInfo{}

type EvaluationInfo struct {
	Id int64
	AudioSrc string
	AudioText string
	TotalScore int64
	OriginalData []byte
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy string
	IsDeleted bool
}
func (m EvaluationInfo) TableName() string {
	return TableEvaluation
}