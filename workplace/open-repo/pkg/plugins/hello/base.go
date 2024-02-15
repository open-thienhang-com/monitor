package hello

import (
	"fmt"

	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
	"mono.thienhang.com/pkg/service"
)

type Hello struct {
	base.Base
}

func NewHello() plugins.Plugin {
	return &Hello{
		Base: base.Base{
			PlugName: "HelloPlugin",
		},
	}
}

func (h *Hello) InitPlugin(srv service.List) {
	app := context.NewApp()
	route := app.Group("hello")
	route.GET("/hello", h.TestHandler)
	h.App = app
}

func (h *Hello) TestHandler(ctx *context.Context) {
	fmt.Println("OKKKKKKKKKK")
	return
}
