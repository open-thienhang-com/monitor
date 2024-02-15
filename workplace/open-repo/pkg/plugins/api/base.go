package api

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
)

type API struct {
	base.Base
}

func NewModel() plugins.Plugin {
	return &API{
		Base: base.Base{
			App:       context.NewApp(),
			PlugName:  "Model Plugin",
			URLPrefix: "Model",
		},
	}
}
