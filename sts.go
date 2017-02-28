package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func usingSTS(account, role string) error {
	sess := session.New(
		&aws.Config{
			Region: aws.String(Config.Region),
		},
	)

	svc := sts.New(sess)

	output, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(fmt.Sprintf("arn:aws:iam::%s:role/%s", account, role)),
		RoleSessionName: aws.String("temp"),
	})
	if err != nil {
		log.Errorf("Unable to assume role: %v", err.Error())
		return err
	}

	accessKey := aws.StringValue(output.Credentials.AccessKeyId)
	secretKey := aws.StringValue(output.Credentials.SecretAccessKey)
	sessionToken := aws.StringValue(output.Credentials.SessionToken)

	os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretKey)
	os.Setenv("AWS_SESSION_TOKEN", sessionToken)

	Config.AccessKey = accessKey
	Config.SecretKey = secretKey
	Config.SessionToken = sessionToken

	return nil
}
