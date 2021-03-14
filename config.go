package aws

import (
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/c3sr/config"
	"github.com/c3sr/vipertags"
)

// AWSRegionUSEast1 ...
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
	done                   chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &awsConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (awsConfig) ConfigName() string {
	return "AWS"
}

// SetDefaults ...
func (a *awsConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *awsConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	a.AccessKey = decrypt(a.AccessKey)
	a.SecretKey = decrypt(a.SecretKey)
	a.SessionToken = decrypt(a.SessionToken)
}

// Wait ...
func (c awsConfig) Wait() {
	<-c.done
}

// String ...
func (c awsConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c awsConfig) Debug() {
	log.Debug("AWS Config = ", c)
}

func init() {
	config.Register(Config)
}
