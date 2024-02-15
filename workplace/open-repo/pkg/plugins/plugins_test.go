package plugins_test

import (
	"testing"

	"mono.thienhang.com/pkg/plugins"
)

func TestLoadFromPlugin(t *testing.T) {
	plugins.LoadFromPlugin("./example/hello.so")
}
