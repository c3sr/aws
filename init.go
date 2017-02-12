package aws

import (
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"

	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

type logwrapper struct {
	*logrus.Entry
}

const (
	debug = true
)

var (
	log         *logwrapper
	awsconf     *aws.Config
	Enabled     = true
	ErrDisabled = errors.New("Amazon package is disabled")
)

func (l *logwrapper) Log(args ...interface{}) {
	log.Debug(args...)
}

func init() {
	config.OnInit(func() {
		log = &logwrapper{
			Entry: logger.New().WithField("pkg", "amazon"),
		}
		logLevel := aws.LogOff
		if config.IsDebug && debug {
			logLevel = aws.LogDebug
		}
		cred := credentials.NewStaticCredentials(
			Config.AccessKey,
			Config.SecretKey,
			"",
		)
	})
}
