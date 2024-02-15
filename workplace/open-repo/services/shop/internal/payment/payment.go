package payment

import "mono.thienhang.com/pkg/models/base"

type Payment struct {
	base.Base
	PaymentMethod string `json:"payment_method"`
}
