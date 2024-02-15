package base

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/database"
	"mono.thienhang.com/pkg/service"
)

type Base struct {
	App       *context.App
	Services  service.List
	Conn      database.Connection
	PlugName  string
	URLPrefix string
	GroupName string
}

func (b *Base) InitPlugin(services service.List) {}

func (b *Base) GetHandler() context.HandlerMap { return b.App.Handlers }
func (b *Base) Name() string                   { return b.PlugName }
func (b *Base) Prefix() string                 { return b.URLPrefix }
func (b *Base) IsInstalled() bool              { return false }
func (b *Base) Uninstall() error               { return nil }
func (b Base) Upgrade() error                  { return nil }

func (b *Base) GetIndexURL() string { return "" }

func (b *Base) InitBase(srv service.List) {
	b.Services = srv
	b.Conn = database.GetConnection(b.Services)
	//
}

func (b *Base) SetPrefix(prefix string) {
	b.URLPrefix = prefix
}
