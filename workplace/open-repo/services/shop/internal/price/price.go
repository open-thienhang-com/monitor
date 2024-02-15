package price

import "mono.thienhang.com/pkg/models/base"

type Price struct {
	base.Base
	Product_ID    int     `json:"product_id"`
	Price_Change  float64 `json:"price_change"`
	Change_status string  `json:"change_status"`
	New_price     float64 `json:"new_price"`
}
