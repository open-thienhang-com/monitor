package plugins

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/service"
)

type Plugin interface {
	GetHandler() context.HandlerMap
	InitPlugin(services service.List)
	Name() string
	Prefix() string
	GetIndexURL() string
	IsInstalled() bool
	Uninstall() error
	Upgrade() error
}
