package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func ConnectAws() *session.Session {
	accessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	secretAccessKey := GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	region := GetEnvWithKey("AWS_REGION")
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				"", // a token will be created when the session it's used.
			),
		},
	)

	if err != nil {
		panic(err)
	}
	return session
}
