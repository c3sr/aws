package aws

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	logger "github.com/c3sr/logger"
)

type logwrapper struct {
	*logrus.Entry
}

var (
	log *logwrapper
)

// Log ...
func (l *logwrapper) Log(args ...interface{}) {
	log.Debug(args...)
}

func init() {
	config.AfterInit(func() {
		log = &logwrapper{
			Entry: logger.New().WithField("pkg", "amazon"),
		}
	})
}
