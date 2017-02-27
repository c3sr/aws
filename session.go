package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

type SessionOptions struct {
	AccessKey string
	SecretKey string
}

type SessionOption func(*SessionOptions)

func AccessKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.AccessKey = s
	}
}

func SecretKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.SecretKey = s
	}
}

func NewSession(opts ...SessionOption) (*session.Session, error) {
	options := SessionOptions{
		AccessKey: Config.AccessKey,
		SecretKey: Config.SecretKey,
	}

	for _, o := range opts {
		o(&options)
	}

	cred := credentials.NewStaticCredentials(
		options.AccessKey,
		options.SecretKey,
		"",
	)

	awsconf := &aws.Config{
		Credentials:      cred,
		Region:           aws.String(Config.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Logger:           log,
	}

	sess, err := session.NewSession(awsconf)
	if err != nil {
		msg := "Was not able to create aws session"
		log.WithError(err).Error(msg)
		err = errors.Wrapf(err, msg)
		return nil, err
	}

	return sess, err
}
