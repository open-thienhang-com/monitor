package inventory

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/cqrs"
)

type Inventory struct {
	cqrs.CQRS
}

func New() plugins.Plugin {
	p := &cqrs.CQRS{}
	p.App = context.NewApp()
	p.Base.PlugName = "Inventory Plugin"
	p.Base.URLPrefix = "internal"
	p.Base.GroupName = "inventory"
	return p
}
