package storageFile










import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"storie/pkg/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)



type AwsS3 struct {

}

func (s AwsS3) UploadTransactions (transactions []domain.Transaction, email string) (string, error) {
	bucketName := os.Getenv("BUCKET_NAME")

	var csvData bytes.Buffer
	writer := csv.NewWriter(&csvData)

	err := writer.Write([]string{"id", "date", "amount"})
	if err != nil {
		return "", err
	}

	for _, t := range transactions {
		amountStr := fmt.Sprintf("%.2f", t.Amount)
		if t.Amount > 0 {
			amountStr = "+" + amountStr
		}

		row := []string{
			fmt.Sprintf("%d", t.Id),
			t.Date,
			amountStr,
		}

		err := writer.Write(row)
		if err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d-%s.csv", time.Now().Unix(), email)



	config := aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("PUBLIC_KEY"), os.Getenv("SECRET"), ""),
	}


	sess, err := session.NewSession(&config)
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(csvData.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	downloadLink := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName)

	return downloadLink, nil
}