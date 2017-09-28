package aws

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// AWSTestSuite ...
type AWSTestSuite struct {
	suite.Suite
}

// SetupTest ...
func (suite *AWSTestSuite) SetupTest() {
}

// TestLoad ...
func (suite *AWSTestSuite) TestLoad() {
	assert.NotNil(suite.T(), Config)
}

// TestPrintable ...
func (suite *AWSTestSuite) TestPrintable() {
	assert.NotEqual(suite.T(), "", Config.String())
}

// TestRegion ...
func (suite *AWSTestSuite) TestRegion() {
	assert.Equal(suite.T(), AWSRegionUSEast1, Config.Region)
}

// TestAWSConfig ...
func TestAWSConfig(t *testing.T) {
	suite.Run(t, new(AWSTestSuite))
}
