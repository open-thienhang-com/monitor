package payment

import "mono.thienhang.com/pkg/service"

func (a *Payment) InitPlugin(srv service.List) {
	a.InitBase(srv)

	// route := a.App.Group(a.GroupName)

	// route.GET("/:id", a.Get)
	// route.POST("/", a.Create)
	// route.PUT("/:id", a.Update)
	// route.DELETE("/:id", a.Delete)
	// route.GET("s", a.GetAll)
	// route.GET("/search", a.Search)

	// //
	// apiv1.HandleFunc("/payment/deposit", c.deposit).Methods(http.MethodGet)
	// //
	// apiv1.HandleFunc("/payment/withdrawl", c.withdrawl).Methods(http.MethodPost)

	// // *** IPN ***
	// apiv1.HandleFunc("/ipn/momo", c.ipnMomo).Methods(http.MethodPost)
	// apiv1.HandleFunc("/ipn/onepay", c.ipnOnePay).Methods(http.MethodGet)
	// apiv1.HandleFunc("/ipn/onepay2", c.ipnOnePay2).Methods(http.MethodGet)
	// apiv1.HandleFunc("/ipn/partner/onepay", c.ipnOnePayPartner).Methods(http.MethodGet)
	// apiv1.HandleFunc("/ipn/zalopay", c.ipnZaloPay).Methods(http.MethodGet)

	//

}
