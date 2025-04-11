package awsService

import (
	"technicaltest/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var AWSSession *session.Session
var AWSS3 *s3.S3

func NewSession() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Conf.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.Conf.AWSAccessKeyID,
			config.Conf.AWSSecretAccessKey,
			"",
		),
	})

	if err != nil {
		panic(err)
	}

	AWSSession = sess
	AWSS3 = s3.New(AWSSession)
}

