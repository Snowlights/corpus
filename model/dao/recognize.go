package dao

import "context"

type RecognizeDao interface {
	AddRecognizeImage(ctx context.Context,data map[string]interface{}) (int64,error)
	AddRecognizeAge(ctx context.Context,data map[string]interface{}) (int64,error)
}
