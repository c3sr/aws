package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	"github.com/c3sr/config"
	"github.com/c3sr/utils"
	"github.com/c3sr/uuid"
)

// SessionOptions ...
type SessionOptions struct {
	AccessKey              string
	SecretKey              string
	SessionToken           string
	Region                 string
	stsRoleDurationSeconds time.Duration
}

// SessionOption ...
type SessionOption func(*SessionOptions)

func decrypt(s string) string {
	if utils.IsEncryptedString(s) {
		c, err := utils.DecryptStringBase64(config.App.Secret, s)
		if err == nil {
			return c
		}
	}
	return s
}

// AccessKey ...
func AccessKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.AccessKey = decrypt(s)
	}
}

// SecretKey ...
func SecretKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.SecretKey = decrypt(s)
	}
}

// Region ...
func Region(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.Region = s
	}
}

// STSRoleDurationSeconds ...
func STSRoleDurationSeconds(t time.Duration) SessionOption {
	return func(opt *SessionOptions) {
		opt.stsRoleDurationSeconds = t
	}
}

// Sts ...
func Sts(data ...string) SessionOption {
	return func(opt *SessionOptions) {
		roleSessionName := uuid.NewV4()
		account := Config.STSAccount
		role := Config.STSRole
		if len(data) >= 1 {
			roleSessionName = data[0]
		}
		if len(data) >= 3 {
			account = data[1]
			role = data[2]
		}
		err := usingSTS(opt, roleSessionName, account, role)
		if err != nil {
			log.WithError(err).Error("Failed to set sts credentials")
		}
	}
}

// NewConfig ...
func NewConfig(opts ...SessionOption) (*aws.Config, error) {
	options := SessionOptions{
		AccessKey:              Config.AccessKey,
		SecretKey:              Config.SecretKey,
		Region:                 Config.Region,
		stsRoleDurationSeconds: Config.STSRoleDurationSeconds,
	}

	for _, o := range opts {
		o(&options)
	}

	cred := credentials.NewStaticCredentials(
		options.AccessKey,
		options.SecretKey,
		options.SessionToken,
	)

	awsconf := &aws.Config{
		Credentials:      cred,
		Region:           aws.String(options.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Logger:           log,
	}

	return awsconf, nil
}

// NewSession ...
func NewSession(opts ...SessionOption) (*session.Session, error) {
	awsconf, err := NewConfig(opts...)
	if err != nil {
		return nil, err
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
