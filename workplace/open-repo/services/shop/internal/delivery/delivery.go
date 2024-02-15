package delivery

import "mono.thienhang.com/pkg/models/base"

type Delivery struct {
	base.Base
	DeliveryMethod string  `json:"delivery_method"`
	Price          float64 `json:"price"`
}
