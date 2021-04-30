package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svcDynamo *dynamodb.DynamoDB

func InitSession() error {
	if svcDynamo == nil {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("eu-west-1")},
		)
		if err != nil {
			svcDynamo = nil
			return err
		}
		svcDynamo = dynamodb.New(sess)
	}
	return nil
}
