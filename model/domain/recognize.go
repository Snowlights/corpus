package domain

const TableRecognize = "corpus_recognize"

var EmptyRecognize = &RecognizeInfo{}

type RecognizeInfo struct {
	Id int64
	AudioSrc string
	RecognizeAgeType string
	ChildScore string
	MiddleScore string
	OldScore string
	GenderType string
	GenderMale string
	GenderFaMale string
	CreatedAt int64
	CreatedBy string
	IsDeleted bool
}
func (m RecognizeInfo) TableName() string {
	return TableRecognize
}

const TablePicture = "corpus_picture"

var EmptyPicture = &PictureInfo{}

type PictureInfo struct {
	Id int64
	PictureSrc string
	PictureDes string
	Md5 string
	PictureText string
	CreatedAt int64
	CreatedBy string
	IsDeleted bool
}
func (m PictureInfo) TableName() string {
	return TablePicture
}