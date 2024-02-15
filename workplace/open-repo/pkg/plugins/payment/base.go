package payment

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
)

const (
	METHOD_MOMO   = "momo"
	METHOD_ZALO   = "zalopay"
	METHOD_ONEPAY = "onepay"
)

type Payment struct {
	base.Base
}

func NewPayment() plugins.Plugin {
	return &Payment{
		Base: base.Base{
			App:      context.NewApp(),
			PlugName: "Payment Plugin",
			// URLPrefix: "Model",
		},
	}
}
