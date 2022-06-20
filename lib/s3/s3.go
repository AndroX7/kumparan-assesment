package s3

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client interface {
	Upload(file io.Reader, path string) error
	Delete(path string) error
}

type S3 struct {
	client  *s3.S3
	session *session.Session
	bucket  *string
}

func NewS3Client() S3Client {
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret := os.Getenv("AWS_ACCESS_KEY_SECRET")
	token := ""

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_S3_BUCKET_REGION")),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecret, token),
	})

	if err != nil {
		log.Println("error when creating session to aws")
	}
	s3Client := s3.New(s)
	return &S3{
		client:  s3Client,
		session: s,
		bucket:  aws.String(os.Getenv("AWS_S3_BUCKET")),
	}
}

func (s *S3) Upload(file io.Reader, path string) error {
	uploader := s3manager.NewUploader(s.session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: s.bucket,
		Key:    aws.String(path),
		Body:   file,
	})
	if err != nil {
		log.Println("error uploading from s3:", err)
		return err
	}

	return nil
}

func (s *S3) Delete(path string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(path),
	})
	if err != nil {
		log.Println("error deleting from s3:", err)
		return err
	}

	return nil
}
