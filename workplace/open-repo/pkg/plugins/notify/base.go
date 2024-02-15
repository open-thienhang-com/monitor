package notify

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins/base"
	"mono.thienhang.com/pkg/plugins/notify/telegram"
	"mono.thienhang.com/pkg/service"
)

type Notify struct {
	base.Base
}

func NewNotify() *Notify {
	return &Notify{
		Base: base.Base{
			PlugName: "NotifyPlugin",
		},
	}
}

func (h *Notify) InitPlugin(srv service.List) {
	h.App = context.NewApp()
	route := h.App.Group("notify")
	//
	route.POST("/telegram", h.TelegramHandler)
	route.POST("/zalo", h.TelegramHandler)
	route.POST("/line", h.TelegramHandler)
	route.POST("/team", h.TelegramHandler)
}

func (h *Notify) TelegramHandler(ctx *context.Context) {
	telegram.SendNotification("XXXXXX")
	return
}
