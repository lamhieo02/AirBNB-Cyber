package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go01-airbnb/config"
	"go01-airbnb/pkg/common"
	"log"
	"net/http"
)

type UploadProvider interface {
	UploadFile(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

type s3Provider struct {
	bucket  string
	region  string
	apiKey  string
	secret  string
	domain  string
	session *session.Session
}

func NewS3Provider(config *config.Config) *s3Provider {
	provider := &s3Provider{
		bucket: config.AWS.S3Bucket,
		region: config.AWS.Region,
		apiKey: config.AWS.APIKey,
		secret: config.AWS.SecretKey,
		domain: config.AWS.S3Domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session
	return provider
}

func (p *s3Provider) UploadFile(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucket),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url: fmt.Sprintf("%s/%s", p.domain, dst),
	}
	return img, nil
}
