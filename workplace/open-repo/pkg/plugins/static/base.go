package static

import (
	"fmt"

	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
	"mono.thienhang.com/pkg/service"
)

type Static struct {
	base.Base
}

func NewStatic() plugins.Plugin {
	return &Static{
		Base: base.Base{
			PlugName: "StaticPlugin",
		},
	}
}

func (h *Static) InitPlugin(srv service.List) {
	app := context.NewApp()
	route := app.Group("Static")
	route.GET("/static", h.TestHandler)
	h.App = app
}

func (h *Static) TestHandler(ctx *context.Context) {
	fmt.Println("OKKKKKKKKKK")
	return
}
