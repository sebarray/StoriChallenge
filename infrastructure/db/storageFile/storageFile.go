package storageFile

import (
	"storie/pkg/domain"
)





type StorageFile interface {
	UploadTransactions(transactions []domain.Transaction , email string) (string, error)
}

var providers = map[string]StorageFile{}

func init() {
	providers["s3"] = AwsS3{}
}

func GetProvider(provider string) StorageFile {
	return providers[provider]
}
