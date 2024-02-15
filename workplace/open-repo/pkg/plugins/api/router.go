package api

import "mono.thienhang.com/pkg/service"

func (a *API) InitPlugin(srv service.List) {
	a.InitBase(srv)

	route := a.App.Group(a.GroupName)

	route.GET("/:id", a.Get)
	route.POST("/", a.Create)
	route.PUT("/:id", a.Update)
	route.DELETE("/:id", a.Delete)
	route.GET("s", a.GetAll)
	route.GET("/search", a.Search)
}
