package aws

import (
	"os"
	"testing"

	"github.com/rai-project/config"
)

func TestMain(m *testing.M) {
	os.Setenv("DEBUG", "TRUE")
	os.Setenv("VERBOSE", "TRUE")
	config.Init()
	config.IsVerbose = true
	config.IsDebug = true
	os.Exit(m.Run())
}
