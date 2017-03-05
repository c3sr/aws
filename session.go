package aws

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
	"github.com/rai-project/uuid"
)

type SessionOptions struct {
	AccessKey    string
	SecretKey    string
	SessionToken string
	Region       string
	EndPoint     string
}

type SessionOption func(*SessionOptions)

func AccessKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.AccessKey = s
	}
}

func EncryptedAccessKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		key := []byte(config.App.Secret)
		sec := []byte(s)
		if p, err := base64.StdEncoding.DecodeString(s); err == nil {
			sec = p
		}
		c, err := utils.Decrypt(key, sec)
		if err != nil {
			log.WithError(err).Error("unable to set encrypted access key")
			return
		}
		opt.AccessKey = string(c)
	}
}

func SecretKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.SecretKey = s
	}
}

func EncryptedSecretKey(s string) SessionOption {
	return func(opt *SessionOptions) {
		key := []byte(config.App.Secret)
		sec := []byte(s)
		if p, err := base64.StdEncoding.DecodeString(s); err == nil {
			sec = p
		}
		c, err := utils.Decrypt(key, sec)
		if err != nil {
			log.WithError(err).Error("unable to set encrypted secret key")
			return
		}
		opt.SecretKey = string(c)
	}
}

func Region(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.Region = s
	}
}

func EndPoint(s string) SessionOption {
	return func(opt *SessionOptions) {
		opt.EndPoint = s
	}
}

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

func NewSession(opts ...SessionOption) (*session.Session, error) {
	options := SessionOptions{
		AccessKey: Config.AccessKey,
		SecretKey: Config.SecretKey,
		Region:    Config.Region,
		EndPoint:  Config.EndPoint,
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
		Endpoint:         aws.String(options.EndPoint),
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
