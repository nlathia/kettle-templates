package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var svcS3 *s3.S3

func InitSession() error {
	if svcS3 == nil {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("eu-west-1")},
		)
		if err != nil {
			svcS3 = nil
			return err
		}
		svcS3 = s3.New(sess)
	}
	return nil
}
