package application

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	GitCommit string
	Version   string
	BuildDate string
)

func GetAppVersion(name string) string {
	return strings.TrimSpace(fmt.Sprintf(
		"%s (%s-%s) (Go %s)",
		name,
		GetVersion(),
		GetCommit(),
		runtime.Version(),
	))
}

func GetVersion() string {
	version := strings.TrimSpace(Version)
	if len(version) == 0 {
		version = "dev"
	}

	return version
}

func GetCommit() string {
	commit := strings.TrimSpace(GitCommit)
	if len(commit) == 0 {
		commit = "dev"
	} else if len(commit) > 8 {
		commit = commit[:8]
	}

	return commit
}

func GetBuildDate() string {
	return BuildDate
}
