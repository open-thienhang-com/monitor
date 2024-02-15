package product

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/api"
)

type Product struct {
	api.API
}

func New() plugins.Plugin {
	p := &Product{}
	p.API.Base.App = context.NewApp()
	p.API.Base.PlugName = "Product Plugin"
	p.API.Base.URLPrefix = "internal"
	p.API.Base.GroupName = "product"
	return p
}

func (plg *Product) Get(ctx *context.Context) {
	p := Model{}
	p.Model.TableName = "product"

	x := p.SetConn(plg.Conn).Find(1)
	api.OkWithDataRaw(ctx, x)
}

func getAll(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("OK2222"))
}

func search(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("OK2222"))
}

func create(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("OK2222"))
}

func update(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("OK2222"))
}

func delete(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("OK2222"))
}
