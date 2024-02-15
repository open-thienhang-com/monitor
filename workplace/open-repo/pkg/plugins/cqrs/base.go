package cqrs

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
	"mono.thienhang.com/pkg/service"
)

type CQRS struct {
	base.Base
}

func NewCQRS() plugins.Plugin {
	return &CQRS{
		Base: base.Base{
			PlugName: "CQRSPlugin",
		},
	}
}

func (h *CQRS) InitPlugin(srv service.List) {
	app := context.NewApp()
	h.App = app
}
