package product

import "mono.thienhang.com/pkg/service"

func (p *Product) InitPlugin(srv service.List) {
	p.InitBase(srv)

	//
	routeAPI := p.App.Group("api/" + p.GroupName)

	routeAPI.GET("/:id", p.Get)
	routeAPI.POST("/", p.Create)
	routeAPI.PUT("/:id", p.Update)
	routeAPI.DELETE("/:id", p.Delete)
	routeAPI.GET("s", p.GetAll)
	routeAPI.GET("/search", p.Search)

	//
	routeView := p.App.Group("views/" + p.GroupName)

	routeView.GET("/:id", p.Get)

}
