package swagger

import (
	"fmt"

	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins/base"
	"mono.thienhang.com/pkg/service"
)

type Swagger struct {
	base.Base
}

func NewSwagger() *Swagger {
	return &Swagger{
		// Base: &plugins.Base{PlugName: "Swagger"},
	}
}

func (e *Swagger) InitPlugin(srv service.List) {
	e.InitBase(srv)
	e.SetPrefix("v1")
	e.App = e.initRouter("", srv)
}

func (e *Swagger) initRouter(prefix string, srv service.List) *context.App {
	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/docs", e.TestHandler)
	route.GET("/swagger", e.JsHandler)
	fmt.Println(app.Handlers)

	return app

}

// func (s *Swagger) GetHandler() context.HandlerMap   { return nil }
// func (s *Swagger) InitPlugin(services service.List) {}
// func (s *Swagger) Name() string                     { return "" }
// func (s *Swagger) Prefix() string                   { return "" }
// func (s *Swagger) GetIndexURL() string              { return "" }
// func (s *Swagger) IsInstalled() bool                { return false }
// func (s *Swagger) Uninstall() error                 { return nil }
// func (s *Swagger) Upgrade() error                   { return nil }

func (e *Swagger) TestHandler(ctx *context.Context) {
	ctx.ServeFile("./docs/index.html", false)
}

func (e *Swagger) JsHandler(ctx *context.Context) {
	ctx.ServeFile("./docs/openapi.js", false)
}
