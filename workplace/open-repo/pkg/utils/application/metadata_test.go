package application

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetVersionDev(t *testing.T) {
	actual := GetVersion()
	require.Equal(t, "dev", actual)
}

func TestGetVersion(t *testing.T) {
	Version = "test"
	actual := GetVersion()
	require.Equal(t, Version, actual)
	Version = ""
}

func TestGetCommitDev(t *testing.T) {
	actual := GetCommit()
	require.Equal(t, "dev", actual)
}

func TestGetCommit(t *testing.T) {
	GitCommit = "test39afdc19206618be53e200453d5b91f4"
	actual := GetCommit()
	require.Equal(t, "test", actual)
	GitCommit = ""
}

func TestGetAppVersionDev(t *testing.T) {
	name := "test name"
	actual := GetAppVersion(name)
	require.Equal(t, fmt.Sprintf("%s (dev-dev) (Go %s)", name, runtime.Version()), actual)
}

func TestGetAppVersion(t *testing.T) {
	Version = "test-version"
	GitCommit = "test39afdc19206618be53e200453d5b91f4"
	name := "test name"
	actual := GetAppVersion(name)
	require.Equal(t, fmt.Sprintf("%s (%s-test) (Go %s)", name, Version, runtime.Version()), actual)
	Version = ""
	GitCommit = ""
}
