package handler

import (
	"os"
	"testing"

	"github.com/shadracnicholas/home-automation/libraries/go/bootstrap"
)

func TestMain(m *testing.M) {
	bootstrap.SetupTest()
	os.Exit(m.Run())
}
