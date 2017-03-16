package aws

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

const (
	// AWSRegionUSEast1 is a helper constant for AWS configs.
	AWSRegionUSEast1 = "us-east-1"
	// AWSRegionUSWest is a helper constant for AWS configs.
	AWSRegionUSWest = "us-west-1"
)

// AWS holds common AWS credentials and keys.
type awsConfig struct {
	AccessKey              string        `json:"access_key" config:"aws.access_key_id" env:"AWS_ACCESS_KEY_ID"`
	SecretKey              string        `json:"secret_key" config:"aws.secret_access_key" env:"AWS_SECRET_ACCESS_KEY"`
	SessionToken           string        `json:"session_token" config:"-" env:"AWS_SESSION_TOKEN"`
	Region                 string        `json:"region" config:"aws.region" default:"us-east-1" env:"AWS_REGION"`
	STSAccount             string        `json:"sts_account" config:"aws.sts_account"`
	STSRole                string        `json:"sts_role" config:"aws.sts_role"`
	STSRoleDurationSeconds time.Duration `json:"sts_role_duration_seconds" config:"aws.sts_role_duration_seconds" default:"30m"` // default is 1 hour
}

var (
	Config = &awsConfig{}
)

func (awsConfig) ConfigName() string {
	return "AWS"
}

func (awsConfig) SetDefaults() {
}

func (a *awsConfig) Read() {
	vipertags.Fill(a)
}

func (c awsConfig) String() string {
	return pp.Sprintln(c)
}

func (c awsConfig) Debug() {
	log.Debug("AWS Config = ", c)
}

func init() {
	config.Register(Config)
}
