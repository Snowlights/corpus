package dao

import "context"

type RecognizeDao interface {
	AddRecognizeImage(ctx context.Context,data map[string]interface{}) (int64,error)
	ListImageByCookie(ctx context.Context,limit,conds map[string]interface{}) ([]string,error)
	CountImageByCookie(ctx context.Context,conds map[string]interface{}) (int64,error)
	AddRecognizeAge(ctx context.Context,data map[string]interface{}) (int64,error)
}
